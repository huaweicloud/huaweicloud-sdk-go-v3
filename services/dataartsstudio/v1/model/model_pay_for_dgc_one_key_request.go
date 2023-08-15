package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PayForDgcOneKeyRequest Request Object
type PayForDgcOneKeyRequest struct {
	Body *OrderReq `json:"body,omitempty"`
}

func (o PayForDgcOneKeyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PayForDgcOneKeyRequest struct{}"
	}

	return strings.Join([]string{"PayForDgcOneKeyRequest", string(data)}, " ")
}
