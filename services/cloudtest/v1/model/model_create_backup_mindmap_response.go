package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateBackupMindmapResponse Response Object
type CreateBackupMindmapResponse struct {
	Code *string `json:"code,omitempty"`

	Data *interface{} `json:"data,omitempty"`

	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateBackupMindmapResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBackupMindmapResponse struct{}"
	}

	return strings.Join([]string{"CreateBackupMindmapResponse", string(data)}, " ")
}
