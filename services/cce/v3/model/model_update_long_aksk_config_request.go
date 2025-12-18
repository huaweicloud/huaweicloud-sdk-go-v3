package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateLongAkskConfigRequest Request Object
type UpdateLongAkskConfigRequest struct {
	Body *UpdateLongAkskConfigRequestBody `json:"body,omitempty"`
}

func (o UpdateLongAkskConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLongAkskConfigRequest struct{}"
	}

	return strings.Join([]string{"UpdateLongAkskConfigRequest", string(data)}, " ")
}
