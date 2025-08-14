package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchListMfaDevicesForUserResponse Response Object
type BatchListMfaDevicesForUserResponse struct {

	// 用户MFA设备信息列表
	UserMfaDevicesEntryList *[]RetrieveMfaDevicesForUserEntryDto `json:"user_mfa_devices_entry_list,omitempty"`
	HttpStatusCode          int                                  `json:"-"`
}

func (o BatchListMfaDevicesForUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchListMfaDevicesForUserResponse struct{}"
	}

	return strings.Join([]string{"BatchListMfaDevicesForUserResponse", string(data)}, " ")
}
