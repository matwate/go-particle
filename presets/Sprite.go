package presets

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	particles "github.com/matwate/go-particles/particles"
)

func SpriteDrawFunc(p *particles.Particle) {

	rl.DrawTextureV(p.Texture, p.Pos, p.Color)
}
