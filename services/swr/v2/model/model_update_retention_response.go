package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateRetentionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateRetentionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRetentionResponse struct{}"
	}

	return strings.Join([]string{"UpdateRetentionResponse", string(data)}, " ")
}
