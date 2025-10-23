package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateMfaDeviceForUserReqBody struct {

	// MFA设备显示名称
	DisplayName string `json:"display_name"`
}

func (o UpdateMfaDeviceForUserReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMfaDeviceForUserReqBody struct{}"
	}

	return strings.Join([]string{"UpdateMfaDeviceForUserReqBody", string(data)}, " ")
}
