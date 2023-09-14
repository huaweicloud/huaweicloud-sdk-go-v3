package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRackRequest Request Object
type ShowRackRequest struct {

	// 机柜ID
	Id string `json:"id"`
}

func (o ShowRackRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRackRequest struct{}"
	}

	return strings.Join([]string{"ShowRackRequest", string(data)}, " ")
}
