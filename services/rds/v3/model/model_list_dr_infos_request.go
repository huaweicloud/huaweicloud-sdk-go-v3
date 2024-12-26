package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDrInfosRequest Request Object
type ListDrInfosRequest struct {

	// 语言。默认en-us。
	XLanguage *string `json:"X-Language,omitempty"`

	Body *QueryDrInfoRequest `json:"body,omitempty"`
}

func (o ListDrInfosRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDrInfosRequest struct{}"
	}

	return strings.Join([]string{"ListDrInfosRequest", string(data)}, " ")
}
