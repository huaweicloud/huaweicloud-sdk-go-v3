package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowPrivateKeyVerifyRequest struct {
	Body *PrivateKeyVerify `json:"body,omitempty"`
}

func (o ShowPrivateKeyVerifyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPrivateKeyVerifyRequest struct{}"
	}

	return strings.Join([]string{"ShowPrivateKeyVerifyRequest", string(data)}, " ")
}
