package internal

type storage struct {
	inMem []UserData
}

func NewStorage() *storage {
	s := &storage{
		inMem: make([]UserData, 0, 10),
	}
	return s
}

func (s *storage) SaveData(data UserData) {
	s.inMem = append(s.inMem, data)
}

func (s *storage) GetAllData() []UserData {
	return s.inMem
}
