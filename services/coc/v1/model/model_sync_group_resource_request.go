package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SyncGroupResourceRequest Request Object
type SyncGroupResourceRequest struct {

	// 分组ID。
	Id string `json:"id"`

	Body *SyncResourceReq `json:"body,omitempty"`
}

func (o SyncGroupResourceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SyncGroupResourceRequest struct{}"
	}

	return strings.Join([]string{"SyncGroupResourceRequest", string(data)}, " ")
}
