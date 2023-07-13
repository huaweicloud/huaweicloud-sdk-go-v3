package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ModifyPublicIpRequestBody struct {

	// 操作标识。取值： - BIND，表示绑定弹性公网IP。 - UNBIND，表示解绑弹性公网IP。
	Action string `json:"action"`

	// 弹性公网IP。
	PublicIp *string `json:"public_ip,omitempty"`

	// 弹性公网IP的ID。
	PublicIpId *string `json:"public_ip_id,omitempty"`
}

func (o ModifyPublicIpRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyPublicIpRequestBody struct{}"
	}

	return strings.Join([]string{"ModifyPublicIpRequestBody", string(data)}, " ")
}
