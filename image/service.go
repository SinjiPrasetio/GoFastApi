package image

type Service interface {
	Upload(userId uint, filename string) (Image, error)
	FindByID(id uint) (Image, error)
	Delete(image Image) error
	List(userId uint, limit int, page int, sort string) ([]Image, error)
	Exists(filename string) bool
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Upload(userId uint, filename string) (Image, error) {
	var image Image
	image.UserID = userId
	image.Image = filename
	res, err := s.repository.Create(image)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (s *service) FindByID(id uint) (Image, error) {
	return s.repository.FindByID(id)
}

func (s *service) Delete(image Image) error {
	err := s.repository.Delete(image)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) List(userId uint, limit int, page int, sort string) ([]Image, error) {

	return s.repository.List(userId, limit, page, sort)
}

func (s *service) Exists(filename string) bool {
	return s.repository.Exists(filename)
}
