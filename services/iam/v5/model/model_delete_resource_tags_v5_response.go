package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteResourceTagsV5Response Response Object
type DeleteResourceTagsV5Response struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteResourceTagsV5Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteResourceTagsV5Response struct{}"
	}

	return strings.Join([]string{"DeleteResourceTagsV5Response", string(data)}, " ")
}
