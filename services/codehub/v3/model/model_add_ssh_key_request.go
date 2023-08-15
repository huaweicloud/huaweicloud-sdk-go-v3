package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddSshKeyRequest Request Object
type AddSshKeyRequest struct {
	Body *AddSshKeyRequestBody `json:"body,omitempty"`
}

func (o AddSshKeyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddSshKeyRequest struct{}"
	}

	return strings.Join([]string{"AddSshKeyRequest", string(data)}, " ")
}
