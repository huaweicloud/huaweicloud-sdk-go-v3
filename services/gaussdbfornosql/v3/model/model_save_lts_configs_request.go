package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SaveLtsConfigsRequest Request Object
type SaveLtsConfigsRequest struct {

	// 语言。
	XLanguage *string `json:"X-Language,omitempty"`

	Body *SaveLtsConfigsRequestBody `json:"body,omitempty"`
}

func (o SaveLtsConfigsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SaveLtsConfigsRequest struct{}"
	}

	return strings.Join([]string{"SaveLtsConfigsRequest", string(data)}, " ")
}
