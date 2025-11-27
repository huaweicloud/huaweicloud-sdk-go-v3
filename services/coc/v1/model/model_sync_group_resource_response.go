package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SyncGroupResourceResponse Response Object
type SyncGroupResourceResponse struct {

	// 同步资源。
	Data           *int64 `json:"data,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o SyncGroupResourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SyncGroupResourceResponse struct{}"
	}

	return strings.Join([]string{"SyncGroupResourceResponse", string(data)}, " ")
}
