package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateStreamForbiddenRequest Request Object
type UpdateStreamForbiddenRequest struct {
	Body *StreamForbiddenSetting `json:"body,omitempty"`
}

func (o UpdateStreamForbiddenRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateStreamForbiddenRequest struct{}"
	}

	return strings.Join([]string{"UpdateStreamForbiddenRequest", string(data)}, " ")
}
