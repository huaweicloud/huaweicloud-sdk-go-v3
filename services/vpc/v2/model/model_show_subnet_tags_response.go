package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSubnetTagsResponse Response Object
type ShowSubnetTagsResponse struct {

	// tag对象列表
	Tags           *[]ResourceTag `json:"tags,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowSubnetTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSubnetTagsResponse struct{}"
	}

	return strings.Join([]string{"ShowSubnetTagsResponse", string(data)}, " ")
}
