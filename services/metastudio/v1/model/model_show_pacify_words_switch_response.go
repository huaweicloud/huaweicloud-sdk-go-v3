package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPacifyWordsSwitchResponse Response Object
type ShowPacifyWordsSwitchResponse struct {

	// 安抚话术功能开关。
	EnablePacifyWords *bool `json:"enable_pacify_words,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowPacifyWordsSwitchResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPacifyWordsSwitchResponse struct{}"
	}

	return strings.Join([]string{"ShowPacifyWordsSwitchResponse", string(data)}, " ")
}
