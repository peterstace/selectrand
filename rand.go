package selectrand

var (
	c1 = make(chan struct{})
	c2 = make(chan struct{})
)

func init() {
	close(c1)
	close(c2)
}

type SelectSource struct{}

func (SelectSource) Int63() int64 {
	var r uint64
	var m uint64 = 1
	for i := 0; i < 63; i++ {
		select {
		case <-c1:
			r = r | m
		case <-c2:
		}
		m *= 2
	}
	return int64(r)
}

func (SelectSource) Seed(int64) {}
