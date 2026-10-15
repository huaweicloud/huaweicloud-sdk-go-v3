package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeBigKeyHotKeySwitchResponse Response Object
type ChangeBigKeyHotKeySwitchResponse struct {

	// 是否成功
	Success *bool `json:"success,omitempty"`

	// 状态值。1-成功，2-失败无需重试，3-失败需要重试
	Status *int32 `json:"status,omitempty"`

	// 错误信息，仅在status=2时返回
	ErrorMsg       *string `json:"error_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ChangeBigKeyHotKeySwitchResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeBigKeyHotKeySwitchResponse struct{}"
	}

	return strings.Join([]string{"ChangeBigKeyHotKeySwitchResponse", string(data)}, " ")
}
