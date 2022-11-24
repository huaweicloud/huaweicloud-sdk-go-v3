package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ConfirmTargetDbTypeRequest struct {
	Body *ConfirmTargetDbReq `json:"body,omitempty"`
}

func (o ConfirmTargetDbTypeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfirmTargetDbTypeRequest struct{}"
	}

	return strings.Join([]string{"ConfirmTargetDbTypeRequest", string(data)}, " ")
}
