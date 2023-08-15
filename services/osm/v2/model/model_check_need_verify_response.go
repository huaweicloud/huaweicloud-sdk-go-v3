package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckNeedVerifyResponse Response Object
type CheckNeedVerifyResponse struct {

	// 是否需要验证
	NeedVerifyCode *int32 `json:"need_verify_code,omitempty"`

	// 过期时间
	ExpireTime     *sdktime.SdkTime `json:"expire_time,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o CheckNeedVerifyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckNeedVerifyResponse struct{}"
	}

	return strings.Join([]string{"CheckNeedVerifyResponse", string(data)}, " ")
}
