package template

import "log"

type transactionRepo interface {
	Find(req Request) (*Response, error)
}

type TemplateService struct {
	repo transactionRepo
}

func NewTemplateService(repo transactionRepo) *TemplateService {
	return &TemplateService{repo: repo}
}

func (s *TemplateService) Process(req Request) (*Response, error) {
	resp, err := s.repo.Find(req)
	if err != nil {
		return nil, err
	}

	log.Println("response: ", resp)

	return resp, nil
}
