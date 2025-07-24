package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateIrackTagsResponse Response Object
type BatchCreateIrackTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchCreateIrackTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateIrackTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateIrackTagsResponse", string(data)}, " ")
}
