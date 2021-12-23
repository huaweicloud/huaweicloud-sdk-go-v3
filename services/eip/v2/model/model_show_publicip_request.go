package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowPublicipRequest struct {
	// 弹性公网IP唯一标识

	PublicipId string `json:"publicip_id"`
}

func (o ShowPublicipRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPublicipRequest struct{}"
	}

	return strings.Join([]string{"ShowPublicipRequest", string(data)}, " ")
}
