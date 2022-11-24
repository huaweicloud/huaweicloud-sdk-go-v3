package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 批量创建保护实例请求数据接口
type BatchCreateProtectedInstancesRequestBody struct {
	ProtectedInstances *BatchCreateProtectedInstancesRequestParams `json:"protected_instances"`
}

func (o BatchCreateProtectedInstancesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateProtectedInstancesRequestBody struct{}"
	}

	return strings.Join([]string{"BatchCreateProtectedInstancesRequestBody", string(data)}, " ")
}
