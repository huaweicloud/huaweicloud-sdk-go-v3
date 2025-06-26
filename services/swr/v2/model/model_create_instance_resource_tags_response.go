package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateInstanceResourceTagsResponse Response Object
type CreateInstanceResourceTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateInstanceResourceTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceResourceTagsResponse struct{}"
	}

	return strings.Join([]string{"CreateInstanceResourceTagsResponse", string(data)}, " ")
}
