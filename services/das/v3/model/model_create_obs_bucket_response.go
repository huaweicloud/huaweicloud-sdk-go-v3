package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateObsBucketResponse Response Object
type CreateObsBucketResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateObsBucketResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateObsBucketResponse struct{}"
	}

	return strings.Join([]string{"CreateObsBucketResponse", string(data)}, " ")
}
