package selectrand

var (
	c1 = make(chan struct{})
	c2 = make(chan struct{})
)

func init() {
	close(c1)
	close(c2)
}

type Source struct{}

func (s Source) Int63() int64 {
	return int64(s.Uint64() >> 1)
}

func (Source) Uint64() uint64 {
	var r uint64
	var m uint64 = 1
	for i := 0; i < 64; i++ {
		select {
		case <-c1:
			r = r | m
		case <-c2:
		}
		m *= 2
	}
	return r
}

func (Source) Seed(int64) {}
