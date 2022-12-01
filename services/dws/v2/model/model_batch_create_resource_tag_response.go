package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchCreateResourceTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchCreateResourceTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateResourceTagResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateResourceTagResponse", string(data)}, " ")
}
