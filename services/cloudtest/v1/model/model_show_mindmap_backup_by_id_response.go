package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMindmapBackupByIdResponse Response Object
type ShowMindmapBackupByIdResponse struct {
	Code *string `json:"code,omitempty"`

	Data *interface{} `json:"data,omitempty"`

	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowMindmapBackupByIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMindmapBackupByIdResponse struct{}"
	}

	return strings.Join([]string{"ShowMindmapBackupByIdResponse", string(data)}, " ")
}
