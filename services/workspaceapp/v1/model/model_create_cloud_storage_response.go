package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCloudStorageResponse Response Object
type CreateCloudStorageResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateCloudStorageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCloudStorageResponse struct{}"
	}

	return strings.Join([]string{"CreateCloudStorageResponse", string(data)}, " ")
}
