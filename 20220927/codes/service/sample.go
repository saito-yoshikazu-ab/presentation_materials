package service

import (
	"20220927/codes/repository"
	"errors"
	"strings"
	"time"
)

type SampleService struct {
	SampleRepository repository.IFSampleRepository
	// s                string
	// gStarRepo        *repository.SampleRepository
	// gRepo            repository.SampleRepository
}

func NewSampleService(sr *repository.SampleRepository) *SampleService {
	return &SampleService{
		SampleRepository: sr,
	}
}

func (s *SampleService) Get(i int) (string, error) {
	// X *repository.SampleRepositoryではなく
	// ○ *repository.MockIFSampleRepositoryのGetNameが実行される
	return s.SampleRepository.GetName(i)
}

func (s *SampleService) Update(i int, name string) (string, error) {
	if !strings.Contains(name, "sample") {
		return "", errors.New("sample repositoryに保存しているnameには必ずsampleが含まれる想定です")
	}
	lastSaveTime, err := s.SampleRepository.GetLastSaveTime(i)
	if err != nil {
		return "", nil
	}
	now := time.Now()
	startDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	if lastSaveTime.After(startDay) {
		return "", errors.New("今日すでに更新済みなら更新できません")
	}

	if err = s.SampleRepository.Update(i, name); err != nil {
		return "", err
	}
	return name, nil
}
