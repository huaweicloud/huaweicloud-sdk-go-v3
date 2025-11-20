package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestoreMetadataResponse Response Object
type RestoreMetadataResponse struct {

	// 任务id
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RestoreMetadataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreMetadataResponse struct{}"
	}

	return strings.Join([]string{"RestoreMetadataResponse", string(data)}, " ")
}
