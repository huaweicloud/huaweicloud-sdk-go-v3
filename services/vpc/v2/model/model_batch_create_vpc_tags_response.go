package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchCreateVpcTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchCreateVpcTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateVpcTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateVpcTagsResponse", string(data)}, " ")
}
