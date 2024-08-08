package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLatestAttachedServerAppResponse Response Object
type ListLatestAttachedServerAppResponse struct {

	// 分发软件信息列表。
	Items          *[]AttachServerAppInfo `json:"items,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ListLatestAttachedServerAppResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLatestAttachedServerAppResponse struct{}"
	}

	return strings.Join([]string{"ListLatestAttachedServerAppResponse", string(data)}, " ")
}
