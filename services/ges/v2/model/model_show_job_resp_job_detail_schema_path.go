package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowJobRespJobDetailSchemaPath struct {

	// OBS文件路径。
	Path *string `json:"path,omitempty"`

	// OBS文件导入日志存储文件。
	Log *string `json:"log,omitempty"`

	// OBS文件导入状态。  - success：完全成功 - failed：完全失败 - partFailed：部分成功
	Status *string `json:"status,omitempty"`

	// 导入失败原因。
	Cause *string `json:"cause,omitempty"`

	// 导入总行数。其值为-1时表示当前版本没有返回该字段。
	TotalLines *int64 `json:"total_lines,omitempty"`

	// 导入失败行数。其值为-1时表示当前版本没有返回该字段。
	FailedLines *int64 `json:"failed_lines,omitempty"`

	// 导入成功行数。其值为-1时表示当前版本没有返回该字段。
	SuccessfulLines *int64 `json:"successful_lines,omitempty"`
}

func (o ShowJobRespJobDetailSchemaPath) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobRespJobDetailSchemaPath struct{}"
	}

	return strings.Join([]string{"ShowJobRespJobDetailSchemaPath", string(data)}, " ")
}
