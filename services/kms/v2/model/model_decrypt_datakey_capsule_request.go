package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DecryptDatakeyCapsuleRequest Request Object
type DecryptDatakeyCapsuleRequest struct {
	Body *DecryptDatakeyCapsuleRequestBody `json:"body,omitempty"`
}

func (o DecryptDatakeyCapsuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DecryptDatakeyCapsuleRequest struct{}"
	}

	return strings.Join([]string{"DecryptDatakeyCapsuleRequest", string(data)}, " ")
}
