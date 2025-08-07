package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowResourceTagsResponse Response Object
type ShowResourceTagsResponse struct {

	// 资源标签
	Tags           *[]ResourceTag `json:"tags,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowResourceTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResourceTagsResponse struct{}"
	}

	return strings.Join([]string{"ShowResourceTagsResponse", string(data)}, " ")
}
