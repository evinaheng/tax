package controller

import "github.com/tax/internal/repository"

// NewSystem New system service
func NewSystem(systemRP repository.System) System {

	svc := &systemSvc{
		systemRP: systemRP,
	}

	return svc
}

func (s *systemSvc) LogOpenFile() {
	s.systemRP.LogOpenFile()
}
