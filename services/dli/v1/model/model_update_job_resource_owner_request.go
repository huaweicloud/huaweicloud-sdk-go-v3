package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateJobResourceOwnerRequest Request Object
type UpdateJobResourceOwnerRequest struct {
	Body *UpdateJobResourceOwnerRequestBody `json:"body,omitempty"`
}

func (o UpdateJobResourceOwnerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateJobResourceOwnerRequest struct{}"
	}

	return strings.Join([]string{"UpdateJobResourceOwnerRequest", string(data)}, " ")
}
