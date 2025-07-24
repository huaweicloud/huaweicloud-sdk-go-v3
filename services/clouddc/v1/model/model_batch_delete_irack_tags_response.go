package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteIrackTagsResponse Response Object
type BatchDeleteIrackTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteIrackTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteIrackTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteIrackTagsResponse", string(data)}, " ")
}
