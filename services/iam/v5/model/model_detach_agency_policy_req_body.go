package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DetachAgencyPolicyReqBody contains information about a urn of a entity
type DetachAgencyPolicyReqBody struct {

	// 委托或信任委托ID，长度为1到64个字符，只包含字母、数字和\"-\"的字符串。
	AgencyId string `json:"agency_id"`
}

func (o DetachAgencyPolicyReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetachAgencyPolicyReqBody struct{}"
	}

	return strings.Join([]string{"DetachAgencyPolicyReqBody", string(data)}, " ")
}
