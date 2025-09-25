package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateP2cVgwRequest Request Object
type CreateP2cVgwRequest struct {

	// 幂等性标识
	XClientToken *string `json:"X-Client-Token,omitempty"`

	Body *CreateP2cVgwRequestBody `json:"body,omitempty"`
}

func (o CreateP2cVgwRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateP2cVgwRequest struct{}"
	}

	return strings.Join([]string{"CreateP2cVgwRequest", string(data)}, " ")
}
