package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeployApplicationRequestBody 应用部署版本信息
type DeployApplicationRequestBody struct {

	// 版本信息。
	Version string `json:"version"`
}

func (o DeployApplicationRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeployApplicationRequestBody struct{}"
	}

	return strings.Join([]string{"DeployApplicationRequestBody", string(data)}, " ")
}
