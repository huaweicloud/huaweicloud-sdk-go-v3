package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateElasticResourcePoolRequestBody 更新弹性资源池请求
type UpdateElasticResourcePoolRequestBody struct {

	// 描述信息。长度限制：256个字符以内。
	Description *string `json:"description,omitempty"`

	// max_cu大于等于该弹性资源池下任意一个队列的最大CU。标准版弹性资源池最小值为64，最大值为32000；基础版弹性资源池最小值为16，最大值为64。
	MaxCu *int32 `json:"max_cu,omitempty"`

	// min_cu大于等于该弹性资源池下所有队列最小CU之和，且小于等于max_cu。标准版弹性资源池最小值为64，最大值为32000；基础版弹性资源池最小值为16，最大值为64。
	MinCu *int32 `json:"min_cu,omitempty"`
}

func (o UpdateElasticResourcePoolRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateElasticResourcePoolRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateElasticResourcePoolRequestBody", string(data)}, " ")
}
