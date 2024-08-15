package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SyncResourceRequest Request Object
type SyncResourceRequest struct {
	Body *SyncResourceReq `json:"body,omitempty"`
}

func (o SyncResourceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SyncResourceRequest struct{}"
	}

	return strings.Join([]string{"SyncResourceRequest", string(data)}, " ")
}
