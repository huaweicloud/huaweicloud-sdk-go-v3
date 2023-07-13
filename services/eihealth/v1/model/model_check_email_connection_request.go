package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckEmailConnectionRequest Request Object
type CheckEmailConnectionRequest struct {
	Body *TryEmailConnectionReq `json:"body,omitempty"`
}

func (o CheckEmailConnectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckEmailConnectionRequest struct{}"
	}

	return strings.Join([]string{"CheckEmailConnectionRequest", string(data)}, " ")
}
