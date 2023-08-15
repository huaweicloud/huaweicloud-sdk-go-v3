package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAuthorisationRequest Request Object
type UpdateAuthorisationRequest struct {

	// 授权实例ID。
	Id string `json:"id"`

	Body *UpdateAuthorisationRequestBody `json:"body,omitempty"`
}

func (o UpdateAuthorisationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAuthorisationRequest struct{}"
	}

	return strings.Join([]string{"UpdateAuthorisationRequest", string(data)}, " ")
}
