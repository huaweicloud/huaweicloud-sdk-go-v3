package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAiAppProtectStatusResponse Response Object
type UpdateAiAppProtectStatusResponse struct {

	// AI应用ID
	AppId *string `json:"appId,omitempty"`

	// 防护状态（0：禁用，1：启用）
	ProtectStatus  *int32 `json:"protect_status,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o UpdateAiAppProtectStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAiAppProtectStatusResponse struct{}"
	}

	return strings.Join([]string{"UpdateAiAppProtectStatusResponse", string(data)}, " ")
}
