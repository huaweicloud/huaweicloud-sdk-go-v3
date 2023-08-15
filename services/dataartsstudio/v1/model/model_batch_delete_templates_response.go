package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteTemplatesResponse Response Object
type BatchDeleteTemplatesResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteTemplatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteTemplatesResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteTemplatesResponse", string(data)}, " ")
}
