package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSharedTagsResponse Response Object
type ShowSharedTagsResponse struct {

	// tag标签的列表。
	Tags           *[]ResourceTag `json:"tags,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowSharedTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSharedTagsResponse struct{}"
	}

	return strings.Join([]string{"ShowSharedTagsResponse", string(data)}, " ")
}
