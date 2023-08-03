package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListConformancePackComplianceByPackIdRequest Request Object
type ListConformancePackComplianceByPackIdRequest struct {

	// 合规规则包ID。
	ConformancePackId string `json:"conformance_pack_id"`

	// 最大的返回数量
	Limit *int32 `json:"limit,omitempty"`

	// 分页参数，通过上一个请求中返回的marker信息作为输入，获取当前页
	Marker *string `json:"marker,omitempty"`

	// 合规规则名称。
	PolicyAssignmentName *string `json:"policy_assignment_name,omitempty"`
}

func (o ListConformancePackComplianceByPackIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConformancePackComplianceByPackIdRequest struct{}"
	}

	return strings.Join([]string{"ListConformancePackComplianceByPackIdRequest", string(data)}, " ")
}
