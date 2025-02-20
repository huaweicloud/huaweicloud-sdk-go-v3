package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowUserDetailRequest Request Object
type ShowUserDetailRequest struct {

	// 用户id。
	UserId string `json:"user_id"`
}

func (o ShowUserDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowUserDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowUserDetailRequest", string(data)}, " ")
}
