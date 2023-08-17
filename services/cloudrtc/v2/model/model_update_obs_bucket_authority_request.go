package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateObsBucketAuthorityRequest Request Object
type UpdateObsBucketAuthorityRequest struct {
	Body *ObsAuthorityConfig `json:"body,omitempty"`
}

func (o UpdateObsBucketAuthorityRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateObsBucketAuthorityRequest struct{}"
	}

	return strings.Join([]string{"UpdateObsBucketAuthorityRequest", string(data)}, " ")
}
