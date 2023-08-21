package cpu

type CPU struct {
	Register registers
}

func New() CPU {
	c := CPU{}
	return c
}
