# vhandler

A dragonfly library to add multiple player and world handler

# examples

<a href="/examples/movement_debug.go">movement_debug.go</a>:

```go
func main() {
	log := logrus.New()
	log.Formatter = &logrus.TextFormatter{ForceColors: true}
	log.Level = logrus.DebugLevel

	cfg := server.DefaultConfig()
	srvCfg, err := cfg.Config(log)
	if err != nil {
		log.Fatal(err)
	}
	srv := srvCfg.New()
	srv.Listen()
	for {
		srv.Accept(func(p *player.Player) {
			setupHandler(p).Set(p)
		})
	}
}

func setupHandler(p *player.Player) *vhandler.Player {
	v := vhandler.NewPlayer()

	v.OnMove(priority.Normal, func(ctx *event.Context, newPos mgl64.Vec3, newYaw, newPitch float64) {
		p.SendTip(fmt.Sprintf("X: %.2f Y: %.2f Z: %.2f\nYaw: %.0f Pitch: %.0f", newPos.X(), newPos.Y(), newPos.Z(), newYaw, newPitch))
	})

	return v
}
```