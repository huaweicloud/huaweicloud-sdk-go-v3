package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IamBpMfaUnconfiguredDetails 未绑定MFA分析详细结果。
type IamBpMfaUnconfiguredDetails struct {

	// 用户的唯一标识符（ID）。
	UserId string `json:"user_id"`

	// 用户的创建时间。
	UserCreatedAt *sdktime.SdkTime `json:"user_created_at"`
}

func (o IamBpMfaUnconfiguredDetails) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IamBpMfaUnconfiguredDetails struct{}"
	}

	return strings.Join([]string{"IamBpMfaUnconfiguredDetails", string(data)}, " ")
}
