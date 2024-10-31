package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMindmapBackupByPageResponse Response Object
type ShowMindmapBackupByPageResponse struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	Data *BasePageInfoMindmapBackup `json:"data,omitempty"`

	// 错误信息
	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowMindmapBackupByPageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMindmapBackupByPageResponse struct{}"
	}

	return strings.Join([]string{"ShowMindmapBackupByPageResponse", string(data)}, " ")
}
