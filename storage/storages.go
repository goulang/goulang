package storage

type FinalStorage interface {
	//上传Token
	GetUploadToken() map[string]interface{}
	FileInfo(string) FinalFileInfo
}

type FinalFileInfo struct {
	FileInfo interface{}
}
