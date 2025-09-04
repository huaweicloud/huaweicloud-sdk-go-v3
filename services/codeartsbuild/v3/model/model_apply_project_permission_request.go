package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ApplyProjectPermissionRequest Request Object
type ApplyProjectPermissionRequest struct {
	Body *ProjectPermissionRequestBody `json:"body,omitempty"`
}

func (o ApplyProjectPermissionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplyProjectPermissionRequest struct{}"
	}

	return strings.Join([]string{"ApplyProjectPermissionRequest", string(data)}, " ")
}
