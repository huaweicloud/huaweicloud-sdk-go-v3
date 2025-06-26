package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteInstanceResourceTagsResponse Response Object
type DeleteInstanceResourceTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteInstanceResourceTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteInstanceResourceTagsResponse struct{}"
	}

	return strings.Join([]string{"DeleteInstanceResourceTagsResponse", string(data)}, " ")
}
