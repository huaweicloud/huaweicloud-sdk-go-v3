package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateOrUpdateStoragePolicyStatementRequest Request Object
type CreateOrUpdateStoragePolicyStatementRequest struct {
	Body *CreateOrUpdateStoragePolicyStatementReq `json:"body,omitempty"`
}

func (o CreateOrUpdateStoragePolicyStatementRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateOrUpdateStoragePolicyStatementRequest struct{}"
	}

	return strings.Join([]string{"CreateOrUpdateStoragePolicyStatementRequest", string(data)}, " ")
}
