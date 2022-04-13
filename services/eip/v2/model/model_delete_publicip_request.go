package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeletePublicipRequest struct {
	// 弹性公网IP唯一标识

	PublicipId string `json:"publicip_id"`
}

func (o DeletePublicipRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePublicipRequest struct{}"
	}

	return strings.Join([]string{"DeletePublicipRequest", string(data)}, " ")
}
