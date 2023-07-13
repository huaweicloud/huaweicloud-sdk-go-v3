package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDedicatedHostTagsResponse Response Object
type ShowDedicatedHostTagsResponse struct {

	// 专属主机标签列表。
	Tags           *[]ResourceTag `json:"tags,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowDedicatedHostTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDedicatedHostTagsResponse struct{}"
	}

	return strings.Join([]string{"ShowDedicatedHostTagsResponse", string(data)}, " ")
}
