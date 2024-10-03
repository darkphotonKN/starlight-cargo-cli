package fileservice

type FileService struct{}

func NewFileService() *FileService {
	return &FileService{}
}

// NOTE: Needs implementation
func (fs *FileService) UploadFile(payload []byte) {
}

// NOTE: Needs implementation
func (fs *FileService) DownloadFile(payload string) ([]byte, error) {
	return []byte{}, nil
}
