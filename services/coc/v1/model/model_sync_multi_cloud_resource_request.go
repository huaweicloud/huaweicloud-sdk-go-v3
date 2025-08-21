package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SyncMultiCloudResourceRequest Request Object
type SyncMultiCloudResourceRequest struct {
	Body *SyncMultiCloudResourceRequestBody `json:"body,omitempty"`
}

func (o SyncMultiCloudResourceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SyncMultiCloudResourceRequest struct{}"
	}

	return strings.Join([]string{"SyncMultiCloudResourceRequest", string(data)}, " ")
}
