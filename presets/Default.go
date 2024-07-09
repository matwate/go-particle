package presets

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
	particles "github.com/matwate/go-particles/particles"
)

func DefaultEmitFunc() particles.Particle {
	return particles.Particle{
		Pos: rl.NewVector2(float32(rl.GetScreenWidth())/2, float32(rl.GetScreenHeight())/2),
		Vel: rl.NewVector2(rand.Float32()*2-1, rand.Float32()*2-1),
		Color: rl.NewColor(
			uint8(rand.Intn(255)),
			uint8(rand.Intn(255)),
			uint8(rand.Intn(255)),
			255,
		),
		Size: 5,
		Life: 1,
	}
}

func DefaultUpdateFunc(p *particles.Particle, dt float32) {
	p.Pos.X += p.Vel.X
	p.Pos.Y += p.Vel.Y

	p.Life -= dt
}

func DefaultDrawFunc(p *particles.Particle) {
	rl.DrawRectangleV(p.Pos, rl.NewVector2(p.Size, p.Size), p.Color)
}
