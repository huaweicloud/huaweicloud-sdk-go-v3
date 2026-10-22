package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetAiAppSettingResponse Response Object
type SetAiAppSettingResponse struct {
	ApiSign        *ApiInfo `json:"api_sign,omitempty"`
	HttpStatusCode int      `json:"-"`
}

func (o SetAiAppSettingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetAiAppSettingResponse struct{}"
	}

	return strings.Join([]string{"SetAiAppSettingResponse", string(data)}, " ")
}
