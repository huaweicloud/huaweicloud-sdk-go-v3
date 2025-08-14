package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteCloudStorageResponse Response Object
type DeleteCloudStorageResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteCloudStorageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCloudStorageResponse struct{}"
	}

	return strings.Join([]string{"DeleteCloudStorageResponse", string(data)}, " ")
}
