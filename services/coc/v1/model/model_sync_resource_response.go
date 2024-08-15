package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SyncResourceResponse Response Object
type SyncResourceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SyncResourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SyncResourceResponse struct{}"
	}

	return strings.Join([]string{"SyncResourceResponse", string(data)}, " ")
}
