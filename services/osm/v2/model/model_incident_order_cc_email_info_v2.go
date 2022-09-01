package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IncidentOrderCcEmailInfoV2 struct {

	// 用户id
	UserId *string `json:"user_id,omitempty" xml:"user_id"`

	// 客户id
	CustomerId *string `json:"customer_id,omitempty" xml:"customer_id"`

	// 操作邮箱
	CcEmail *string `json:"cc_email,omitempty" xml:"cc_email"`
}

func (o IncidentOrderCcEmailInfoV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IncidentOrderCcEmailInfoV2 struct{}"
	}

	return strings.Join([]string{"IncidentOrderCcEmailInfoV2", string(data)}, " ")
}
