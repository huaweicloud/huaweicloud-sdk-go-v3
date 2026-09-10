package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttachShareFilesystemRequest Request Object
type AttachShareFilesystemRequest struct {
	Body *AttachShareFilesystemRequestBody `json:"body,omitempty"`
}

func (o AttachShareFilesystemRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachShareFilesystemRequest struct{}"
	}

	return strings.Join([]string{"AttachShareFilesystemRequest", string(data)}, " ")
}
