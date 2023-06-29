package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResetLoginMethodRequest Request Object
type ResetLoginMethodRequest struct {

	// 堡垒机实例ID，使用UUID格式。
	ServerId string `json:"server_id"`
}

func (o ResetLoginMethodRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetLoginMethodRequest struct{}"
	}

	return strings.Join([]string{"ResetLoginMethodRequest", string(data)}, " ")
}
