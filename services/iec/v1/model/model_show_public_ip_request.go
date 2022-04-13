package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowPublicIpRequest struct {
	// 弹性公网IP ID。

	PublicipId string `json:"publicip_id"`
}

func (o ShowPublicIpRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPublicIpRequest struct{}"
	}

	return strings.Join([]string{"ShowPublicIpRequest", string(data)}, " ")
}
