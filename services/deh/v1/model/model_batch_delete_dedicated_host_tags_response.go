package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchDeleteDedicatedHostTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteDedicatedHostTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteDedicatedHostTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteDedicatedHostTagsResponse", string(data)}, " ")
}
