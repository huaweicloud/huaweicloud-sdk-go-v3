package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateShareConnectionsRequest struct {

	// 语言
	XLanguage *string `json:"X-Language,omitempty"`

	Body *CreateShareConnectionsRequestBody `json:"body,omitempty"`
}

func (o CreateShareConnectionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateShareConnectionsRequest struct{}"
	}

	return strings.Join([]string{"CreateShareConnectionsRequest", string(data)}, " ")
}
