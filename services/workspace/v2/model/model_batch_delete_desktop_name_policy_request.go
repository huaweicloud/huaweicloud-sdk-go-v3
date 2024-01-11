package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteDesktopNamePolicyRequest Request Object
type BatchDeleteDesktopNamePolicyRequest struct {
	Body *BatchDeleteDesktopNamePolicyReq `json:"body,omitempty"`
}

func (o BatchDeleteDesktopNamePolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteDesktopNamePolicyRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteDesktopNamePolicyRequest", string(data)}, " ")
}
