package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateNoticeRuleReq struct {

	// API版本，固定值“v1”，该值不可修改。
	ApiVersion string `json:"api_version"`

	// API类型，固定值“NoticeRule”，该值不可修改。
	Kind string `json:"kind"`

	Spec *UpdateNoticeRuleItem `json:"spec"`
}

func (o UpdateNoticeRuleReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNoticeRuleReq struct{}"
	}

	return strings.Join([]string{"UpdateNoticeRuleReq", string(data)}, " ")
}
