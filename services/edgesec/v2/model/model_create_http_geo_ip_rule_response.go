package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateHttpGeoIpRuleResponse Response Object
type CreateHttpGeoIpRuleResponse struct {

	// 规则id
	Id *string `json:"id,omitempty"`

	// 规则名称
	Name *string `json:"name,omitempty"`

	// 规则所在策略id
	Policyid *string `json:"policyid,omitempty"`

	// 规则所在策略名称
	PolicyName *string `json:"policy_name,omitempty"`

	// 创建规则时间戳
	Timestamp *int64 `json:"timestamp,omitempty"`

	// 规则描述
	Description *string `json:"description,omitempty"`

	// 规则开关状态
	Status *int32 `json:"status,omitempty"`

	// 地理位置
	GeoIp *string `json:"geo_ip,omitempty"`

	// 地理位置列表
	GeoTagList *[]string `json:"geo_tag_list,omitempty"`

	// 拦截/放行/仅记录
	White          *int32 `json:"white,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CreateHttpGeoIpRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHttpGeoIpRuleResponse struct{}"
	}

	return strings.Join([]string{"CreateHttpGeoIpRuleResponse", string(data)}, " ")
}
