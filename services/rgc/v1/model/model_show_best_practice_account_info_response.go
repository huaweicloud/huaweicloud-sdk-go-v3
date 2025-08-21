package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBestPracticeAccountInfoResponse Response Object
type ShowBestPracticeAccountInfoResponse struct {

	// 账号类型
	AccountType *string `json:"account_type,omitempty"`

	// 有效开始时间
	EffectiveStartTime *sdktime.SdkTime `json:"effective_start_time,omitempty"`

	// 有效过期时间
	EffectiveExpirationTime *sdktime.SdkTime `json:"effective_expiration_time,omitempty"`

	// 当前时间
	CurrentTime    *sdktime.SdkTime `json:"current_time,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ShowBestPracticeAccountInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBestPracticeAccountInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowBestPracticeAccountInfoResponse", string(data)}, " ")
}
