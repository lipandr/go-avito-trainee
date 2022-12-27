package internal

type service struct {
	stg *storage
}

func NewService(stg *storage) *service {
	s := &service{
		stg: stg,
	}
	return s
}

func (s *service) saveData(data UserData) {
	s.stg.SaveData(data)
}

func (s *service) getAllData() []UserData {
	return s.stg.GetAllData()
}
