package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateAppAuthorizationsRequest Request Object
type BatchUpdateAppAuthorizationsRequest struct {
	Body *BatchAssignAppAuthorizationsReq `json:"body,omitempty"`
}

func (o BatchUpdateAppAuthorizationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateAppAuthorizationsRequest struct{}"
	}

	return strings.Join([]string{"BatchUpdateAppAuthorizationsRequest", string(data)}, " ")
}
