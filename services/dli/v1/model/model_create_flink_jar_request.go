package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateFlinkJarRequest Request Object
type CreateFlinkJarRequest struct {
	Body *CreateFlinkJarRequestBody `json:"body,omitempty"`
}

func (o CreateFlinkJarRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFlinkJarRequest struct{}"
	}

	return strings.Join([]string{"CreateFlinkJarRequest", string(data)}, " ")
}
