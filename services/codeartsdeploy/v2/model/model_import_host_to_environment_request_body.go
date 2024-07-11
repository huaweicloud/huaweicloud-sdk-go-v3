package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportHostToEnvironmentRequestBody 环境下导入主机请求体
type ImportHostToEnvironmentRequestBody struct {

	// 主机集群id
	GroupId string `json:"group_id"`

	// 要导入的主机id列表
	HostIds []string `json:"host_ids"`
}

func (o ImportHostToEnvironmentRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportHostToEnvironmentRequestBody struct{}"
	}

	return strings.Join([]string{"ImportHostToEnvironmentRequestBody", string(data)}, " ")
}
