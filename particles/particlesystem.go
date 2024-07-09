package particle

type ParticleEmitter struct {
	Particles    []Particle
	MaxParticles int
	EmissionRate float32
	EmitFunc     func() Particle
	UpdateFunc   func(*Particle, float32)
	DrawFunc     func(*Particle)
}

// NewEmitter creates a new ParticleEmitter with the specified parameters.
//
// Parameters:
// - maxParticles: the maximum number of particles that can be emitted.
// - emissionRate: the rate at which particles are emitted.
// - emitFunc: a function that returns a new Particle.
// - updateFunc: a function that updates a Particle based on the elapsed time.
// - drawFunc: a function that draws a Particle.
//
// Returns:
// - *ParticleEmitter: a pointer to the newly created ParticleEmitter.
func NewEmitter(maxParticles int, emissionRate float32, emitFunc func() Particle, updateFunc func(*Particle, float32), drawFunc func(*Particle)) *ParticleEmitter {

	return &ParticleEmitter{
		Particles:    make([]Particle, 0, maxParticles),
		MaxParticles: maxParticles,
		EmissionRate: emissionRate,
		EmitFunc:     emitFunc,
		UpdateFunc:   updateFunc,
		DrawFunc:     drawFunc,
	}
}

// Emit adds a new particle to the particle emitter if the maximum number of particles has not been reached.
//
// No parameters.
// No return value.
func (em *ParticleEmitter) Emit() {
	if len(em.Particles) < em.MaxParticles {
		em.Particles = append(em.Particles, em.EmitFunc())
	}
}

// Update updates the particles in the ParticleEmitter based on the elapsed time.
//
// Parameters:
// - dt: the elapsed time since the last update.
//
// No return value.
func (em *ParticleEmitter) Update(dt float32) {
	for i := range em.Particles {
		var p *Particle
		if i < len(em.Particles) {
			p = &em.Particles[i]
		} else {
			continue
		}

		em.UpdateFunc(p, dt)
		if p.Life <= 0 {
			em.Particles = append(em.Particles[:i], em.Particles[i+1:]...)

			i--
		}

	}
}

// Draw executes the DrawFunc for each particle in the ParticleEmitter.
//
// No parameters.
// No return value.
func (em *ParticleEmitter) Draw() {

	for _, p := range em.Particles {

		em.DrawFunc(&p)
	}
}
