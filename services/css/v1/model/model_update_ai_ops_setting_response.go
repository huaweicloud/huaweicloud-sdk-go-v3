package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAiOpsSettingResponse Response Object
type UpdateAiOpsSettingResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateAiOpsSettingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAiOpsSettingResponse struct{}"
	}

	return strings.Join([]string{"UpdateAiOpsSettingResponse", string(data)}, " ")
}
