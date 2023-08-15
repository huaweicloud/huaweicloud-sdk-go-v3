package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAuthInfoRequest Request Object
type UpdateAuthInfoRequest struct {
	Body *UpdateAuthInfoRequestBody `json:"body,omitempty"`
}

func (o UpdateAuthInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAuthInfoRequest struct{}"
	}

	return strings.Join([]string{"UpdateAuthInfoRequest", string(data)}, " ")
}
