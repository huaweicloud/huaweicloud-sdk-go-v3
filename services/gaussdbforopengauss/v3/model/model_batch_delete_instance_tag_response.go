package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteInstanceTagResponse Response Object
type BatchDeleteInstanceTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteInstanceTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteInstanceTagResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteInstanceTagResponse", string(data)}, " ")
}
