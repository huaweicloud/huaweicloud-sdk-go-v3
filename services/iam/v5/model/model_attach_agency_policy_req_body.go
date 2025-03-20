package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttachAgencyPolicyReqBody Contains information about a id of an agency.
type AttachAgencyPolicyReqBody struct {

	// 委托或信任委托ID，长度为1到64个字符，只包含字母、数字和\"-\"的字符串。
	AgencyId string `json:"agency_id"`
}

func (o AttachAgencyPolicyReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachAgencyPolicyReqBody struct{}"
	}

	return strings.Join([]string{"AttachAgencyPolicyReqBody", string(data)}, " ")
}
