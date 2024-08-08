package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateImageServerRequest Request Object
type CreateImageServerRequest struct {

	// CBC接口回调时，请求头里带上的业务ID。
	ServiceTransactionId *string `json:"Service-Transaction-Id,omitempty"`

	Body *CreateImageServerReq `json:"body,omitempty"`
}

func (o CreateImageServerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateImageServerRequest struct{}"
	}

	return strings.Join([]string{"CreateImageServerRequest", string(data)}, " ")
}
