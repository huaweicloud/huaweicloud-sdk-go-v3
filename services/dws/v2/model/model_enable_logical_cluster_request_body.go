package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnableLogicalClusterRequestBody 切换逻辑集群开关请求
type EnableLogicalClusterRequestBody struct {

	// true-切换开关
	Enable bool `json:"enable"`
}

func (o EnableLogicalClusterRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableLogicalClusterRequestBody struct{}"
	}

	return strings.Join([]string{"EnableLogicalClusterRequestBody", string(data)}, " ")
}
