package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListFreeResourceUsagesRequest struct {
	// 语言。中文：zh_CN英文：en_US缺省为zh_CN。

	XLanguage *string `json:"X-Language,omitempty"`

	Body *ListFreeResourceUsagesReq `json:"body,omitempty"`
}

func (o ListFreeResourceUsagesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFreeResourceUsagesRequest struct{}"
	}

	return strings.Join([]string{"ListFreeResourceUsagesRequest", string(data)}, " ")
}
