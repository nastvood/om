package conf

type (
	config struct {
		Capacity    int
		Concurrency bool
	}

	Option func(conf *config)
)

func DefaultConfig() config {
	return config{
		Concurrency: true,
	}
}

func WithCapacity(capacity int) Option {
	return func(conf *config) {
		conf.Capacity = capacity
	}
}

func WithoutConcurrency() Option {
	return func(conf *config) {
		conf.Concurrency = false
	}
}
