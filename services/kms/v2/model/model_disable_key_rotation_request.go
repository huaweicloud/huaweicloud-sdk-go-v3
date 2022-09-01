package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DisableKeyRotationRequest struct {
	Body *OperateKeyRequestBody `json:"body,omitempty" xml:"body"`
}

func (o DisableKeyRotationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisableKeyRotationRequest struct{}"
	}

	return strings.Join([]string{"DisableKeyRotationRequest", string(data)}, " ")
}
