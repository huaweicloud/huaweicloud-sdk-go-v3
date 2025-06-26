package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSubResourceTagsResponse Response Object
type DeleteSubResourceTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteSubResourceTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSubResourceTagsResponse struct{}"
	}

	return strings.Join([]string{"DeleteSubResourceTagsResponse", string(data)}, " ")
}
