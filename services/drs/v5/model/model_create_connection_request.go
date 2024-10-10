package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateConnectionRequest Request Object
type CreateConnectionRequest struct {

	// 请求语言类型。
	XLanguage *string `json:"X-Language,omitempty"`

	Body *CreateConnectionReq `json:"body,omitempty"`
}

func (o CreateConnectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateConnectionRequest struct{}"
	}

	return strings.Join([]string{"CreateConnectionRequest", string(data)}, " ")
}
