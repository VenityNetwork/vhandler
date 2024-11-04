package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func fetchURL(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to make GET request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 response: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return body, nil
}

type HandlerArg struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type Handler struct {
	FuncName string       `json:"funcName"`
	Comments []string     `json:"comments"`
	Args     []HandlerArg `json:"args"`
}

func (h Handler) Cancellable() bool {
	for _, arg := range h.Args {
		if arg.Type == "*event.Context" {
			return true
		}
	}
	return false
}

func main() {
	handlers, err := parseHandler("https://raw.githubusercontent.com/df-mc/dragonfly/refs/heads/master/server/player/handler.go")
	if err != nil {
		panic(fmt.Errorf("failed to parse player handler: %w", err))
	}

	worldHandlers, err := parseHandler("https://raw.githubusercontent.com/df-mc/dragonfly/refs/heads/master/server/world/handler.go")
	if err != nil {
		panic(fmt.Errorf("failed to parse world handler: %w", err))
	}

	//_ = os.WriteFile("world_handler.go", []byte(generateImpl("World", "*world.World", "w", worldHandlers)), 0644)
	//_ = os.WriteFile("player_handler.go", []byte(generateImpl("Player", "*player.Player", "p", handlers)), 0644)
	_ = os.WriteFile("player_native_bridge_handler.go", []byte(generateBridge("Player", "player.Handler", "*player.Player", "p", handlers)), 0644)
	_ = os.WriteFile("world_native_bridge_handler.go", []byte(generateBridge("World", "world.Handler", "*world.World", "w", worldHandlers)), 0644)
	_ = os.WriteFile("player_handler_func_types.go", []byte(generateFuncTypes("Player", "*player.Player", "p", handlers)), 0644)
	_ = os.WriteFile("world_handler_func_types.go", []byte(generateFuncTypes("World", "*world.World", "w", worldHandlers)), 0644)
	_ = os.WriteFile("player_handlers.go", []byte(generateHandlerList("Player", handlers)), 0644)
	_ = os.WriteFile("world_handlers.go", []byte(generateHandlerList("World", worldHandlers)), 0644)
}

func generateFuncTypes(prefix string, param1 string, param1Name string, handlers []Handler) string {
	str := strings.Builder{}
	str.WriteString("package vhandler\n\n")
	for _, handler := range handlers {
		str.WriteString("type ")
		str.WriteString(prefix)
		str.WriteString(handler.FuncName)
		str.WriteString("Func func(")
		str.WriteString(param1Name)
		str.WriteString(" ")
		str.WriteString(param1)
		if len(handler.Args) > 0 {
			str.WriteString(", ")
		}
		for i, arg := range handler.Args {
			str.WriteString(arg.Name)
			str.WriteString(" ")
			str.WriteString(arg.Type)
			if i != len(handler.Args)-1 {
				str.WriteString(", ")
			}
		}
		str.WriteString(")\n")
	}
	return str.String()
}

func generateHandlerList(name string, handlers []Handler) string {
	str := strings.Builder{}
	str.WriteString("package vhandler\n\n")

	str.WriteString("type ")
	str.WriteString(name)
	str.WriteString("Handlers struct {\n")
	for _, handler := range handlers {
		fn := handler.FuncName[len("Handle"):]
		fn = strings.ToLower(fn[:1]) + fn[1:]
		str.WriteString("\t")
		str.WriteString(fn)
		str.WriteString("Handlers []*handlerWrapper[")
		str.WriteString(name)
		str.WriteString(handler.FuncName)
		str.WriteString("Func]\n")
	}
	str.WriteString("}\n\n")

	str.WriteString("func New")
	str.WriteString(name)
	str.WriteString("Handlers() *")
	str.WriteString(name)
	str.WriteString("Handlers {\n")
	str.WriteString("\treturn &")
	str.WriteString(name)
	str.WriteString("Handlers{\n")
	for _, handler := range handlers {
		fn := handler.FuncName[len("Handle"):]
		fn = strings.ToLower(fn[:1]) + fn[1:]
		str.WriteString("\t\t")
		str.WriteString(fn)
		str.WriteString("Handlers: make([]*handlerWrapper[")
		str.WriteString(name)
		str.WriteString(handler.FuncName)
		str.WriteString("Func], 0),\n")
	}
	str.WriteString("\t}\n")
	str.WriteString("}\n\n")

	// implement Add method
	for _, handler := range handlers {
		fn2 := handler.FuncName[len("Handle"):]
		fn := strings.ToLower(fn2[:1]) + fn2[1:]
		str.WriteString("func (h *")
		str.WriteString(name)
		str.WriteString("Handlers) On")
		str.WriteString(fn2)
		str.WriteString("(handler ")
		str.WriteString(name)
		str.WriteString(handler.FuncName)
		str.WriteString("Func, priority Priority) {\n")
		str.WriteString("\th.")
		str.WriteString(fn)
		str.WriteString("Handlers = append(h.")
		str.WriteString(fn)
		str.WriteString("Handlers, &handlerWrapper[")
		str.WriteString(name)
		str.WriteString(handler.FuncName)
		str.WriteString("Func]{priority, handler})\n")
		str.WriteString("\tsortHandlers(h.")
		str.WriteString(fn)
		str.WriteString("Handlers)\n")

		str.WriteString("}\n\n")
	}

	return str.String()
}

func generateBridge(name string, nativeHandler string, param1 string, param1Name string, handlers []Handler) string {
	str := strings.Builder{}
	str.WriteString("package vhandler\n\n")

	str.WriteString("type ")
	str.WriteString(name)
	str.WriteString("NativeBridgeHandler struct {\n")
	str.WriteString("\t")
	str.WriteString(param1Name)
	str.WriteString(" ")
	str.WriteString(param1)
	str.WriteString("\n")
	str.WriteString("\th *")
	str.WriteString(name)
	str.WriteString("Handlers\n")
	str.WriteString("}\n\n")

	str.WriteString("// Compile-time check to ensure that the ")
	str.WriteString(name)
	str.WriteString("NativeBridgeHandler implements the ")
	str.WriteString(name)
	str.WriteString("Handler interface.\n")
	str.WriteString("var _ ")
	str.WriteString(nativeHandler)
	str.WriteString(" = &")
	str.WriteString(name)
	str.WriteString("NativeBridgeHandler{}\n\n")

	for _, handler := range handlers {
		str.WriteString("func (h *")
		str.WriteString(name)
		str.WriteString("NativeBridgeHandler) ")
		str.WriteString(handler.FuncName)
		str.WriteString("(")
		for i, arg := range handler.Args {
			str.WriteString(arg.Name)
			str.WriteString(" ")
			str.WriteString(arg.Type)
			if i != len(handler.Args)-1 {
				str.WriteString(", ")
			}
		}
		str.WriteString(") {\n")

		str.WriteString("\tfor _, handler := range h.h.")
		fn := handler.FuncName[len("Handle"):]
		fn = strings.ToLower(fn[:1]) + fn[1:]
		str.WriteString(fn)
		str.WriteString("Handlers {\n")
		str.WriteString("\t\thandler.h")
		str.WriteString("(h.")
		str.WriteString(param1Name)
		if len(handler.Args) > 0 {
			str.WriteString(", ")
		}
		for i, arg := range handler.Args {
			str.WriteString(arg.Name)
			if i != len(handler.Args)-1 {
				str.WriteString(", ")
			}
		}
		str.WriteString(")\n")
		str.WriteString("\t}\n")
		str.WriteString("}\n\n")
	}

	return str.String()
}

//func generateImpl(name string, param1 string, param1Name string, handlers []Handler) string {
//	str := strings.Builder{}
//	str.WriteString("package vhandler\n\n")
//
//	str.WriteString("type ")
//	str.WriteString(name)
//	str.WriteString("Handler interface {\n")
//	for _, handler := range handlers {
//		str.WriteString("\t// ")
//		str.WriteString(strings.Join(handler.Comments, "\n\t// "))
//		str.WriteString("\n")
//		str.WriteString("\t")
//		str.WriteString(handler.FuncName)
//		str.WriteString("(")
//		str.WriteString(param1Name)
//		str.WriteString(" ")
//		str.WriteString(param1)
//		if len(handler.Args) > 0 {
//			str.WriteString(", ")
//		}
//		for i, arg := range handler.Args {
//			str.WriteString(arg.Name)
//			str.WriteString(" ")
//			str.WriteString(arg.Type)
//			if i != len(handler.Args)-1 {
//				str.WriteString(", ")
//			}
//		}
//		str.WriteString(")\n")
//	}
//	str.WriteString("}\n\n")
//
//	str.WriteString("type ")
//	str.WriteString(name)
//	str.WriteString("NopHandler struct {}\n\n")
//	for _, handler := range handlers {
//		str.WriteString(fmt.Sprintf("func (nop %sNopHandler) ", name))
//		str.WriteString(handler.FuncName)
//		str.WriteString("(")
//		str.WriteString(param1)
//		if len(handler.Args) > 0 {
//			str.WriteString(", ")
//		}
//		for i, arg := range handler.Args {
//			str.WriteString(arg.Type)
//			if i != len(handler.Args)-1 {
//				str.WriteString(", ")
//			}
//		}
//		str.WriteString(") { \n\t// nop!\n}\n\n")
//	}
//	return str.String()
//}

func parseHandler(url string) ([]Handler, error) {
	content, err := fetchURL(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch URL: %w", err)
	}

	lines := strings.Split(string(content), "\n")
	scan := false

	handlers := make([]Handler, 0)

	curComments := make([]string, 0)

	for _, line := range lines {
		if line == "type Handler interface {" {
			scan = true
			continue
		}
		if !scan {
			continue
		}
		if strings.HasPrefix(line, "\t// ") {
			curComments = append(curComments, line[4:])
		}
		if strings.HasPrefix(line, "\tHandle") {
			line = line[1:]
			funcName := strings.Split(line, "(")[0]
			args := strings.Split(strings.Split(line, "(")[1], ")")[0]
			argList := strings.Split(args, ",")
			handlerArgs := make([]HandlerArg, 0)
			for _, arg := range argList {
				arg = strings.TrimSpace(arg)
				if arg == "" {
					continue
				}
				parts := strings.Split(arg, " ")
				hasType := len(parts) == 2
				argType := ""
				if hasType {
					argType = parts[1]
				}

				switch argType {
				case "World", "Sound", "Liquid", "Block", "Entity":
					argType = "world." + argType
				}

				handlerArgs = append(handlerArgs, HandlerArg{
					Name: parts[0],
					Type: argType,
				})
			}

			hasArgs := len(handlerArgs) > 0
			if hasArgs {
				curType := handlerArgs[len(handlerArgs)-1].Type
				for i := len(handlerArgs) - 1; i >= 0; i-- {
					if handlerArgs[i].Type == "" {
						handlerArgs[i].Type = curType
					} else {
						curType = handlerArgs[i].Type
					}
				}
			}

			handlers = append(handlers, Handler{
				FuncName: funcName,
				Args:     handlerArgs,
				Comments: curComments,
			})
			curComments = make([]string, 0)
		}
		if strings.HasPrefix(line, "}") {
			break
		}
	}

	return handlers, nil
}
