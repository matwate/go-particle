package particle

import rl "github.com/gen2brain/raylib-go/raylib"

type Particle struct {
	Pos     rl.Vector2
	Vel     rl.Vector2
	Color   rl.Color
	Texture rl.Texture2D
	Size    float32
	Life    float32
}
