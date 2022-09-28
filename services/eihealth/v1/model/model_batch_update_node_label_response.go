package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchUpdateNodeLabelResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchUpdateNodeLabelResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateNodeLabelResponse struct{}"
	}

	return strings.Join([]string{"BatchUpdateNodeLabelResponse", string(data)}, " ")
}
