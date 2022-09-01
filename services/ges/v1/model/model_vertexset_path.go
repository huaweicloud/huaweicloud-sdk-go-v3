package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 点数据
type VertexsetPath struct {

	// OBS文件路径
	Path string `json:"path" xml:"path"`

	// OBS文件导入日志存储文件
	Log *string `json:"log,omitempty" xml:"log"`

	// - OBS文件导入状态。 - success：完全成功 - failed：完全失败 - partFailed：部分成功
	Status string `json:"status" xml:"status"`

	// 导入失败原因
	Cause *string `json:"cause,omitempty" xml:"cause"`

	// 导入总行数。其值为-1时表示当前版本没有返回该字段。
	TotalLines *int64 `json:"totalLines,omitempty" xml:"totalLines"`

	// 导入失败行数。其值为-1时表示当前版本没有返回该字段。
	FailedLines *int64 `json:"failedLines,omitempty" xml:"failedLines"`

	// 导出成功行数。其值为-1时表示当前版本没有返回该字段。
	SuccessfulLines *int64 `json:"successfulLines,omitempty" xml:"successfulLines"`
}

func (o VertexsetPath) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VertexsetPath struct{}"
	}

	return strings.Join([]string{"VertexsetPath", string(data)}, " ")
}
