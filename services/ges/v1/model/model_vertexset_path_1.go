package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VertexsetPath1 struct {
	// 导入OBS文件对应的jobId。

	JobId string `json:"jobId"`
	// OBS存储路径，不包含OBS Endpoint。

	Path string `json:"path"`
	// OBS文件导入状态。  - success：完全导入成功 - partiallyFailed：部分失败 - failed：完全导入失败

	Status string `json:"status"`
}

func (o VertexsetPath1) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VertexsetPath1 struct{}"
	}

	return strings.Join([]string{"VertexsetPath1", string(data)}, " ")
}
