package mathx

// Option sets an optional parameter for the function.
type Option func(cfg *config)

type config struct {
	vector []float64
}

// WithBuf sets float buffer for function.
func WithBuf(buf []float64) Option {
	return func(cfg *config) {
		cfg.vector = buf
	}
}
