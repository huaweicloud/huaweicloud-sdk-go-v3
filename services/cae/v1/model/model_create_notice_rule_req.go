package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateNoticeRuleReq struct {

	// API版本，固定值“v1”，该值不可修改。
	ApiVersion string `json:"api_version"`

	// API类型，固定值“NoticeRule”，该值不可修改。
	Kind string `json:"kind"`

	Spec *CreateNoticeRuleItem `json:"spec"`
}

func (o CreateNoticeRuleReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNoticeRuleReq struct{}"
	}

	return strings.Join([]string{"CreateNoticeRuleReq", string(data)}, " ")
}
