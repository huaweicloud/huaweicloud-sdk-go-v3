package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchRebuildDesktopsSystemDiskRequest Request Object
type BatchRebuildDesktopsSystemDiskRequest struct {
	Body *RebuildDesktopsReq `json:"body,omitempty"`
}

func (o BatchRebuildDesktopsSystemDiskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchRebuildDesktopsSystemDiskRequest struct{}"
	}

	return strings.Join([]string{"BatchRebuildDesktopsSystemDiskRequest", string(data)}, " ")
}
