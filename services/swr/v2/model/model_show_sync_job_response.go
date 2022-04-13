package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowSyncJobResponse struct {
	Body           *[]SyncJob `json:"body,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ShowSyncJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSyncJobResponse struct{}"
	}

	return strings.Join([]string{"ShowSyncJobResponse", string(data)}, " ")
}
