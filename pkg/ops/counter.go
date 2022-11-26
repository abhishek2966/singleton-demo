package ops

type counter struct {
	count int
}

var instance *counter

func GetCounterInstance() *counter {
	if instance == nil {
		instance = new(counter)
	}
	return instance
}
func (s *counter) AddOne() int {
	s.count++
	return s.count
}

func (s *counter) GetCount() (count int) {
	if s == nil {
		return
	}
	return s.count
}
