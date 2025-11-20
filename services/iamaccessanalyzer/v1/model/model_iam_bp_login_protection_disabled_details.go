package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IamBpLoginProtectionDisabledDetails 登录保护未开启分析详细结果。
type IamBpLoginProtectionDisabledDetails struct {

	// 用户的唯一标识符（ID）。
	UserId string `json:"user_id"`

	// 用户的创建时间。
	UserCreatedAt *sdktime.SdkTime `json:"user_created_at"`
}

func (o IamBpLoginProtectionDisabledDetails) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IamBpLoginProtectionDisabledDetails struct{}"
	}

	return strings.Join([]string{"IamBpLoginProtectionDisabledDetails", string(data)}, " ")
}
