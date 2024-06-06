package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateNoticeRuleResponse Response Object
type UpdateNoticeRuleResponse struct {

	// API版本，固定值“v1”，该值不可修改。
	ApiVersion *string `json:"api_version,omitempty"`

	// API类型，固定值“NoticeRule”，该值不可修改。
	Kind *string `json:"kind,omitempty"`

	Spec           *CreateNoticeRuleRespItem `json:"spec,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o UpdateNoticeRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNoticeRuleResponse struct{}"
	}

	return strings.Join([]string{"UpdateNoticeRuleResponse", string(data)}, " ")
}
