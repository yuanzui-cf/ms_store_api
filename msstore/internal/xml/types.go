package xml

// FileInfo 表示文件标识符和修改信息
type FileInfo struct {
	Identifier string
	Modified   string
}

// UpdateIdentity 表示更新ID和修订号
type UpdateIdentity struct {
	UpdateID       string
	RevisionNumber string
}
