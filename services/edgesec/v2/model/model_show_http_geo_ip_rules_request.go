package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHttpGeoIpRulesRequest Request Object
type ShowHttpGeoIpRulesRequest struct {

	// 策略id
	PolicyId string `json:"policy_id"`

	// 分页查询参数，第page页
	Page *int32 `json:"page,omitempty"`

	// 分页查询参数，每页显示pagesize条记录
	Pagesize *int32 `json:"pagesize,omitempty"`
}

func (o ShowHttpGeoIpRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHttpGeoIpRulesRequest struct{}"
	}

	return strings.Join([]string{"ShowHttpGeoIpRulesRequest", string(data)}, " ")
}
