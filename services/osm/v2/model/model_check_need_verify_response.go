package model

import (
	"encoding/json"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"

	"strings"
)

// Response Object
type CheckNeedVerifyResponse struct {
	// 是否需要验证

	NeedVerifyCode *int32 `json:"need_verify_code,omitempty"`
	// 过期时间

	ExpireTime     *sdktime.SdkTime `json:"expire_time,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o CheckNeedVerifyResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CheckNeedVerifyResponse struct{}"
	}

	return strings.Join([]string{"CheckNeedVerifyResponse", string(data)}, " ")
}
