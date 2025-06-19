package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CopyJobRequest Request Object
type CopyJobRequest struct {
	Body *CopyBuildJobRequestBody `json:"body,omitempty"`
}

func (o CopyJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CopyJobRequest struct{}"
	}

	return strings.Join([]string{"CopyJobRequest", string(data)}, " ")
}
