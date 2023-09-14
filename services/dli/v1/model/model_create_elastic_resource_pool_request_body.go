package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateElasticResourcePoolRequestBody 创建弹性资源池消息
type CreateElasticResourcePoolRequestBody struct {

	// 新建的弹性资源池名称，名称只能包含数字、小写英文字母和下划线，但不能是纯数字，且不能以下划线开头。长度限制：1~128个字符。
	ElasticResourcePoolName string `json:"elastic_resource_pool_name"`

	// 描述信息。长度限制：256个字符以内。
	Description *string `json:"description,omitempty"`

	// 虚拟集群关联的vpc cidr.如果不填，默认值为172.16.0.0//12
	CidrInVpc *string `json:"cidr_in_vpc,omitempty"`

	// 最大CU大于等于该资源池下任意一个队列的最大CU之和且大于min_cu。最小值为64
	MaxCu int32 `json:"max_cu"`

	// 计费类型 1、按需计费
	ChargingMode *int32 `json:"charging_mode,omitempty"`

	// 最小CU大于等于该资源池下所有队列最小CU之和,最小值为64
	MinCu int32 `json:"min_cu"`

	// 企业ID，不填默认为“0”
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 标签
	Tags *[]TmsTag `json:"tags,omitempty"`
}

func (o CreateElasticResourcePoolRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateElasticResourcePoolRequestBody struct{}"
	}

	return strings.Join([]string{"CreateElasticResourcePoolRequestBody", string(data)}, " ")
}
