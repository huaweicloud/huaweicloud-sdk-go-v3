package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListConsumeSubCustomersRequest Request Object
type ListConsumeSubCustomersRequest struct {

	// 语言。中文：zh_CN英文：en_US缺省为zh_CN。
	XLanguage *string `json:"X-Language,omitempty"`

	Body *ListConsumeSubCustomersReq `json:"body,omitempty"`
}

func (o ListConsumeSubCustomersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConsumeSubCustomersRequest struct{}"
	}

	return strings.Join([]string{"ListConsumeSubCustomersRequest", string(data)}, " ")
}
