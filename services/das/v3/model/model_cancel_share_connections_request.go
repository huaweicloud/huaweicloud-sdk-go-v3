package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CancelShareConnectionsRequest Request Object
type CancelShareConnectionsRequest struct {

	// 语言
	XLanguage *string `json:"X-Language,omitempty"`

	Body *CancelShareConnectionsRequestBody `json:"body,omitempty"`
}

func (o CancelShareConnectionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelShareConnectionsRequest struct{}"
	}

	return strings.Join([]string{"CancelShareConnectionsRequest", string(data)}, " ")
}
