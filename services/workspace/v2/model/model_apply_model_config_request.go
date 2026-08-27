package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ApplyModelConfigRequest Request Object
type ApplyModelConfigRequest struct {
	Body *ApplyModelConfigReq `json:"body,omitempty"`
}

func (o ApplyModelConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplyModelConfigRequest struct{}"
	}

	return strings.Join([]string{"ApplyModelConfigRequest", string(data)}, " ")
}
