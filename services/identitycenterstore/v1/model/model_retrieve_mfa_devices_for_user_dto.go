package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RetrieveMfaDevicesForUserDto struct {

	// 身份源的全局唯一标识符（ID）
	IdentityStoreId string `json:"identity_store_id"`

	// 用户唯一标识符（ID）
	UserId string `json:"user_id"`
}

func (o RetrieveMfaDevicesForUserDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RetrieveMfaDevicesForUserDto struct{}"
	}

	return strings.Join([]string{"RetrieveMfaDevicesForUserDto", string(data)}, " ")
}
