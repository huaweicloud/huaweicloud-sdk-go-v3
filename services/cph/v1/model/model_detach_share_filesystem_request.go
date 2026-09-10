package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DetachShareFilesystemRequest Request Object
type DetachShareFilesystemRequest struct {
	Body *DetachShareFilesystemRequestBody `json:"body,omitempty"`
}

func (o DetachShareFilesystemRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetachShareFilesystemRequest struct{}"
	}

	return strings.Join([]string{"DetachShareFilesystemRequest", string(data)}, " ")
}
