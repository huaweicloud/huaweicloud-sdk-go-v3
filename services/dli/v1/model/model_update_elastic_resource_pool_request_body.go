package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateElasticResourcePoolRequestBody 更新弹性资源池请求
type UpdateElasticResourcePoolRequestBody struct {

	// 描述信息。长度限制：256个字符以内。
	Description *string `json:"description,omitempty"`

	// 弹性资源池的最大CU数
	MaxCu *int32 `json:"max_cu,omitempty"`

	// 弹性资源池的最小CU数
	MinCu *int32 `json:"min_cu,omitempty"`
}

func (o UpdateElasticResourcePoolRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateElasticResourcePoolRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateElasticResourcePoolRequestBody", string(data)}, " ")
}
