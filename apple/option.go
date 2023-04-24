package apple

type options struct {
	password string
}

func NewOption() *options {
	return &options{}
}

type Option func(opt *options)

func WithPassword(password string) Option {
	return func(opt *options) {
		opt.password = password
	}
}
