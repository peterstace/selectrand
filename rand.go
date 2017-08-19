package selectrand

var (
	c1 = make(chan struct{})
	c2 = make(chan struct{})
)

func init() {
	close(c1)
	close(c2)
}

func b() byte {
	var b byte
	for i := uint(0); i < 8; i++ {
		select {
		case <-c1:
			b = b | (1 << i)
		case <-c2:
		}
	}
	return b
}

type SelectSource struct{}

func (SelectSource) Int63() int64 {
	return 0 // TODO
}

func (SelectSource) Seed(int64) {
}
