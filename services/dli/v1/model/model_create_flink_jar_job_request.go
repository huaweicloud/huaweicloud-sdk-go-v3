package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateFlinkJarJobRequest Request Object
type CreateFlinkJarJobRequest struct {
	Body *CreateFlinkJarJobRequestBody `json:"body,omitempty"`
}

func (o CreateFlinkJarJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFlinkJarJobRequest struct{}"
	}

	return strings.Join([]string{"CreateFlinkJarJobRequest", string(data)}, " ")
}
