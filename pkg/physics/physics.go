package physics

import (
	"ray-tracer/pkg/tuples"
)

type Environment struct {
	Gravity tuples.Tuple
	Wind    tuples.Tuple
}

type Projectile struct {
	Position tuples.Tuple
	Velocity tuples.Tuple
}

func Tick(env Environment, proj Projectile) Projectile {
	position := tuples.Add(proj.Position, proj.Velocity)
	velocity := tuples.Add(tuples.Add(proj.Velocity, env.Gravity), env.Wind)
	return Projectile{position, velocity}
}
