package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHotWordsSwitchResponse Response Object
type ShowHotWordsSwitchResponse struct {

	// 热词功能开关。
	EnableHotWords *bool `json:"enable_hot_words,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowHotWordsSwitchResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHotWordsSwitchResponse struct{}"
	}

	return strings.Join([]string{"ShowHotWordsSwitchResponse", string(data)}, " ")
}
