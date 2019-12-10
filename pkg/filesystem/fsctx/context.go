package fsctx

type key int

const (
	// GinCtx Gin的上下文
	GinCtx key = iota
	// SavePathCtx 文件物理路径
	SavePathCtx
	// FileHeaderCtx 上传的文件
	FileHeaderCtx
	// PathCtx 文件或目录的虚拟路径
	PathCtx
	// FileModelCtx 文件数据库模型
	FileModelCtx
)