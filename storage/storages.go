package storage

type FinalStorage interface {

	PutFile(string, []byte) (FinalPutFile, error)
	DeleteFile(string) error
	BatchDeleteFile([]string) error
	FileInfo(string) FinalFileInfo
	PrefixListFiles(string, int) FinalListItem
	ChangeMimeType(string, string) error
}

type FinalFileInfo struct {
	FileInfo interface{}
}

type FinalListItem struct {
	ListItem interface{}
}

type FinalPutFile struct {
	PutFile interface{}
}
