package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTrustPolicyV5Request Request Object
type UpdateTrustPolicyV5Request struct {

	// 信任委托ID，长度为1到64个字符，只包含字母、数字和\"-\"的字符串。
	AgencyId string `json:"agency_id"`

	Body *UpdateTrustPolicyReqBody `json:"body,omitempty"`
}

func (o UpdateTrustPolicyV5Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTrustPolicyV5Request struct{}"
	}

	return strings.Join([]string{"UpdateTrustPolicyV5Request", string(data)}, " ")
}
