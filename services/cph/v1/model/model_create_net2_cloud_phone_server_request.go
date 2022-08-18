package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateNet2CloudPhoneServerRequest struct {
	Body *CreateNet2CloudPhoneServerRequestBody `json:"body,omitempty"`
}

func (o CreateNet2CloudPhoneServerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNet2CloudPhoneServerRequest struct{}"
	}

	return strings.Join([]string{"CreateNet2CloudPhoneServerRequest", string(data)}, " ")
}
