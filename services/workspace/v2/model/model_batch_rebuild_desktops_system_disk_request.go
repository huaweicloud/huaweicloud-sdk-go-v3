package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchRebuildDesktopsSystemDiskRequest Request Object
type BatchRebuildDesktopsSystemDiskRequest struct {

	// CBC接口回调时，请求头里带上的业务ID
	ServiceTransactionId *string `json:"Service-Transaction-Id,omitempty"`

	Body *RebuildDesktopsReq `json:"body,omitempty"`
}

func (o BatchRebuildDesktopsSystemDiskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchRebuildDesktopsSystemDiskRequest struct{}"
	}

	return strings.Join([]string{"BatchRebuildDesktopsSystemDiskRequest", string(data)}, " ")
}
