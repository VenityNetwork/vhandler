# vhandler

A dragonfly functional event handler.

For now, the documentation is not complete. Please refer to the source code for more information.

## Installation

```bash
go get github.com/venitynetwork/vhandler
```

## Usage

```go
func setupHandler() *vhandler.PlayerHandlers {
    h := vhandler.NewPlayerHandlers()
    h.OnAttackEntity(func (p *player.Player, ctx *event.Context, e world.Entity, force *float64, height *float64, critical *bool) {
        fmt.Printf("%s attacked %v\n", p.Name(), e)
    }, vhandler.PriorityNormal)
    
    h.OnBlockBreak(func (p *player.Player, ctx *event.Context, pos cube.Pos, drops *[]item.Stack, xp *int) {
        fmt.Printf("%s broke a block at %v\n", p.Name(), pos)
    }, vhandler.PriorityLow)
	
    return h
}

func setPlayerHandler(p *player.Player){
    h := setupHandler()
    // Note: h can be reused for multiple players
    vhandler.HandlePlayer(p, h)
}
```