package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 构建工程。
type BuildInfo struct {
	// 构建ID，查看构建列表获取。

	Id *string `json:"id,omitempty"`

	Parameters *BuildInfoParameters `json:"parameters,omitempty"`
}

func (o BuildInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BuildInfo struct{}"
	}

	return strings.Join([]string{"BuildInfo", string(data)}, " ")
}
