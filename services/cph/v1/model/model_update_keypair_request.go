package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateKeypairRequest struct {
	Body *UpdateKeypairRequestBody `json:"body,omitempty" xml:"body"`
}

func (o UpdateKeypairRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateKeypairRequest struct{}"
	}

	return strings.Join([]string{"UpdateKeypairRequest", string(data)}, " ")
}
