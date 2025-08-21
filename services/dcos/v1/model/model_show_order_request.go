package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowOrderRequest Request Object
type ShowOrderRequest struct {

	// 服务单号。
	Number string `json:"number"`
}

func (o ShowOrderRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOrderRequest struct{}"
	}

	return strings.Join([]string{"ShowOrderRequest", string(data)}, " ")
}
