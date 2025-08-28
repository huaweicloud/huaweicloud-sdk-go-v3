package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SaveTtscTenantConfigsRequestBody struct {

	// 主键id
	Id *string `json:"id,omitempty"`

	// 租户级的配置项，当前只用于词表
	Key string `json:"key"`

	// 租户级的配置项的值
	Value string `json:"value"`
}

func (o SaveTtscTenantConfigsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SaveTtscTenantConfigsRequestBody struct{}"
	}

	return strings.Join([]string{"SaveTtscTenantConfigsRequestBody", string(data)}, " ")
}
