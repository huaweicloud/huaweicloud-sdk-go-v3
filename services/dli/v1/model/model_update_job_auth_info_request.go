package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateJobAuthInfoRequest Request Object
type UpdateJobAuthInfoRequest struct {
	Body *UpdateJobAuthInfoRequestBody `json:"body,omitempty"`
}

func (o UpdateJobAuthInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateJobAuthInfoRequest struct{}"
	}

	return strings.Join([]string{"UpdateJobAuthInfoRequest", string(data)}, " ")
}
