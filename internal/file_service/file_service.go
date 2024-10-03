package fileservice

import "fmt"

type FileService struct{}

func NewFileService() *FileService {
	return &FileService{}
}

/**
* Handles file uploads by requesting file path and uploading the file to the connected server.
**/
func (fs *FileService) UploadFile(file []byte) {
	fmt.Println("File:", file)
}

// NOTE: Needs implementation
func (fs *FileService) DownloadFile(payload string) ([]byte, error) {
	return []byte{}, nil
}
