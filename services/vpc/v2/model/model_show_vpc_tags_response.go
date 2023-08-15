package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowVpcTagsResponse Response Object
type ShowVpcTagsResponse struct {

	// tag对象列表
	Tags           *[]ResourceTag `json:"tags,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowVpcTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVpcTagsResponse struct{}"
	}

	return strings.Join([]string{"ShowVpcTagsResponse", string(data)}, " ")
}
