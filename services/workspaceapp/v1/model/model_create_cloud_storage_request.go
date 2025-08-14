package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCloudStorageRequest Request Object
type CreateCloudStorageRequest struct {
	Body *CreateCloudStorageReq `json:"body,omitempty"`
}

func (o CreateCloudStorageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCloudStorageRequest struct{}"
	}

	return strings.Join([]string{"CreateCloudStorageRequest", string(data)}, " ")
}
