package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowIegInfoRequest Request Object
type ShowIegInfoRequest struct {

	// 智能企业网关ID
	IegId string `json:"ieg_id"`
}

func (o ShowIegInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIegInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowIegInfoRequest", string(data)}, " ")
}
