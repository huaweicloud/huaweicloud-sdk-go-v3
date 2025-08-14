package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchListMfaDevicesForUserReqBody struct {

	// 用户列表
	UserList []RetrieveMfaDevicesForUserDto `json:"user_list"`
}

func (o BatchListMfaDevicesForUserReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchListMfaDevicesForUserReqBody struct{}"
	}

	return strings.Join([]string{"BatchListMfaDevicesForUserReqBody", string(data)}, " ")
}
