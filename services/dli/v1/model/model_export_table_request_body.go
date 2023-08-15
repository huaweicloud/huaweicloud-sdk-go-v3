package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportTableRequestBody 导出数据
type ExportTableRequestBody struct {

	// 导出数据的储存路径（当前仅支持将数据存储在OBS上，且OBS路径须以s3a开头）。另外，“export_mode”配置为“errorifexists”时，该路径下的文件夹须不存在，如请求样例中的“test”文件夹。
	DataPath string `json:"data_path"`

	// 导出数据的类型（当前仅支持csv格式数据）。
	DataType string `json:"data_type"`

	// 被导出数据的表所在的数据库名称。
	DatabaseName string `json:"database_name"`

	// 被导出数据的表名称。
	TableName string `json:"table_name"`

	// 导出数据的压缩方法。目前支持gzip、bzip2、deflate压缩方式；若不希望压缩，则输入none。
	Compress string `json:"compress"`

	// 指定执行该任务的队列。若不指定队列，将采用default队列执行操作。
	QueueName *string `json:"queue_name,omitempty"`

	// 导出模式，目前支持“ErrorIfExists”，“Overwrite”，不指定“export_mode”则默认为“ErrorIfExists”。  “ErrorIfExists”：存在即报错。指定的导出目录必须不存在，如果指定目录已经存在，系统将返回错误信息，无法执行导出操作。 “Overwrite”：覆盖。在指定目录下新建文件，会删除已有文件。
	ExportMode *string `json:"export_mode,omitempty"`

	// 导出csv格式数据时，是否导出列名。  设置为“true”，表示导出列名。 设置为“false”，表示不导出列名。 若为空，默认为“false”。
	WithColumnHeader *bool `json:"with_column_header,omitempty"`
}

func (o ExportTableRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportTableRequestBody struct{}"
	}

	return strings.Join([]string{"ExportTableRequestBody", string(data)}, " ")
}
