package error

type Error interface {
	error
	Code() string
}

func Check(err error) {
	if err != nil {
		panic(err)
	}
}
