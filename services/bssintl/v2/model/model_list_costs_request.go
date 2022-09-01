package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListCostsRequest struct {

	// 语言。中文：zh_CN英文：en_US。缺省为zh_CN。
	XLanguage *string `json:"X-Language,omitempty" xml:"X-Language"`

	Body *ListCostsReq `json:"body,omitempty" xml:"body"`
}

func (o ListCostsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCostsRequest struct{}"
	}

	return strings.Join([]string{"ListCostsRequest", string(data)}, " ")
}
