package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateKeypairRequest Request Object
type UpdateKeypairRequest struct {
	Body *UpdateKeypairRequestBody `json:"body,omitempty"`
}

func (o UpdateKeypairRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateKeypairRequest struct{}"
	}

	return strings.Join([]string{"UpdateKeypairRequest", string(data)}, " ")
}
