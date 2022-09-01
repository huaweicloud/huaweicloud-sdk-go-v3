package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListFreeResourceInfosRequest struct {

	// 语言。中文：zh_CN英文：en_US缺省为zh_CN。
	XLanguage *string `json:"X-Language,omitempty" xml:"X-Language"`

	Body *ListFreeResourceInfosReq `json:"body,omitempty" xml:"body"`
}

func (o ListFreeResourceInfosRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFreeResourceInfosRequest struct{}"
	}

	return strings.Join([]string{"ListFreeResourceInfosRequest", string(data)}, " ")
}
