package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CloseAiOpsSettingResponse Response Object
type CloseAiOpsSettingResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CloseAiOpsSettingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CloseAiOpsSettingResponse struct{}"
	}

	return strings.Join([]string{"CloseAiOpsSettingResponse", string(data)}, " ")
}
