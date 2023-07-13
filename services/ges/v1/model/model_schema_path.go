package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SchemaPath 点数据
type SchemaPath struct {

	// OBS文件路径
	Path string `json:"path"`

	// OBS文件导入日志存储文件
	Log *string `json:"log,omitempty"`

	// - OBS文件导入状态。 - success：完全成功 - failed：完全失败 - partFailed：部分成功
	Status string `json:"status"`

	// 导入失败原因
	Cause *string `json:"cause,omitempty"`
}

func (o SchemaPath) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SchemaPath struct{}"
	}

	return strings.Join([]string{"SchemaPath", string(data)}, " ")
}
