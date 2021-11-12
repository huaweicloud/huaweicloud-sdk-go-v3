package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AttachEipRequestBody struct {
	// 公网IP的ID。

	PublicIpId string `json:"public_ip_id"`
	// 公网IP。

	PublicIp string `json:"public_ip"`
}

func (o AttachEipRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachEipRequestBody struct{}"
	}

	return strings.Join([]string{"AttachEipRequestBody", string(data)}, " ")
}
