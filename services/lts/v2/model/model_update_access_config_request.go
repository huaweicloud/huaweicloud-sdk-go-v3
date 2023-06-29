package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAccessConfigRequest Request Object
type UpdateAccessConfigRequest struct {
	Body *UpdateAccessConfigRequestBody `json:"body,omitempty"`
}

func (o UpdateAccessConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAccessConfigRequest struct{}"
	}

	return strings.Join([]string{"UpdateAccessConfigRequest", string(data)}, " ")
}
