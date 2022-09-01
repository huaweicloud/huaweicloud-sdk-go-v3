package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DisableKeyRequest struct {
	Body *OperateKeyRequestBody `json:"body,omitempty" xml:"body"`
}

func (o DisableKeyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisableKeyRequest struct{}"
	}

	return strings.Join([]string{"DisableKeyRequest", string(data)}, " ")
}
