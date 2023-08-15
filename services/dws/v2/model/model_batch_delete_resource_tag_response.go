package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteResourceTagResponse Response Object
type BatchDeleteResourceTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteResourceTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteResourceTagResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteResourceTagResponse", string(data)}, " ")
}
