package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteRetentionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteRetentionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRetentionResponse struct{}"
	}

	return strings.Join([]string{"DeleteRetentionResponse", string(data)}, " ")
}
