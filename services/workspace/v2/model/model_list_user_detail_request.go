package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUserDetailRequest Request Object
type ListUserDetailRequest struct {

	// 用户ID。
	UserId string `json:"user_id"`
}

func (o ListUserDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUserDetailRequest struct{}"
	}

	return strings.Join([]string{"ListUserDetailRequest", string(data)}, " ")
}
