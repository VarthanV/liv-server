package fileservice

type Service struct {
	filesToTrack []string
}

func NewFileService() *Service {
	return &Service{}
}
