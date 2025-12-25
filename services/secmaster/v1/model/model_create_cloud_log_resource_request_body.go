package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateCloudLogResourceRequestBody struct {

	// 租户ID
	DomainId string `json:"domain_id"`

	// 资源列表
	Resources []ResourceDto `json:"resources"`
}

func (o CreateCloudLogResourceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCloudLogResourceRequestBody struct{}"
	}

	return strings.Join([]string{"CreateCloudLogResourceRequestBody", string(data)}, " ")
}
