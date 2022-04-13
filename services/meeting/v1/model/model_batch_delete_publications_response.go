package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchDeletePublicationsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeletePublicationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeletePublicationsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeletePublicationsResponse", string(data)}, " ")
}
