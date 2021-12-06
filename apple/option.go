package apple

type option struct {
	password string
}

func NewOption() *option {
	return &option{}
}

type Option func(opt *option)

func WithPassword(password string) Option {
	return func(opt *option) {
		opt.password = password
	}
}
