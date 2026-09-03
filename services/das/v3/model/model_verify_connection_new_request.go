package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VerifyConnectionNewRequest Request Object
type VerifyConnectionNewRequest struct {
	Body *VerifyConnectionNewRequestBody `json:"body,omitempty"`
}

func (o VerifyConnectionNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VerifyConnectionNewRequest struct{}"
	}

	return strings.Join([]string{"VerifyConnectionNewRequest", string(data)}, " ")
}
