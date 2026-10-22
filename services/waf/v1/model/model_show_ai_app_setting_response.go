package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAiAppSettingResponse Response Object
type ShowAiAppSettingResponse struct {
	ApiSign        *ApiInfo `json:"api_sign,omitempty"`
	HttpStatusCode int      `json:"-"`
}

func (o ShowAiAppSettingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAiAppSettingResponse struct{}"
	}

	return strings.Join([]string{"ShowAiAppSettingResponse", string(data)}, " ")
}
