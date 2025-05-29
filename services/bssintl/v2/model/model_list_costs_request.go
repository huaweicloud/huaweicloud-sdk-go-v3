package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCostsRequest Request Object
type ListCostsRequest struct {

	// 语言，字段预留。默认zh_cn，枚举：zh_cn：中文 en_us：英文
	XLanguage *string `json:"X-Language,omitempty"`

	Body *ListCostsReq `json:"body,omitempty"`
}

func (o ListCostsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCostsRequest struct{}"
	}

	return strings.Join([]string{"ListCostsRequest", string(data)}, " ")
}
