package particles

type Particle struct {
    lifetime int
    speed float64

    x float64
    y float64
}

type updateFunc = func(particle *Particle, deltaMS int64)

type ParticleParams struct {
    Maxlife int64
    MaxSpeed float64

    ParticleCount int

    X float64
    Y float64
}

type ParticleSystem struct{
    ParticleParams

    lastTime int64
    place func(particle *Particle, deltaMS int64)
}


func NewParticleSystem(params ParticleParams, updateFunc) ParticleSystem {

}

