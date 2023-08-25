package conf

type (
	Config struct {
		Capacity    int
		Concurrency bool
		BucketLen   int
	}

	Option func(conf *Config)
)

func DefaultConfig() Config {
	return Config{
		Concurrency: true,
	}
}

func WithCapacity(capacity int) Option {
	return func(conf *Config) {
		conf.Capacity = capacity
	}
}

func WithoutConcurrency() Option {
	return func(conf *Config) {
		conf.Concurrency = false
	}
}

func WithBucketLen(bucketLen int) Option {
	return func(conf *Config) {
		conf.BucketLen = bucketLen
	}
}
