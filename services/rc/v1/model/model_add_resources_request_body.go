package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddResourcesRequestBody struct {

	// 资源标识列表
	ResourceIds []string `json:"resource_ids"`
}

func (o AddResourcesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddResourcesRequestBody struct{}"
	}

	return strings.Join([]string{"AddResourcesRequestBody", string(data)}, " ")
}
