package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResourceSpecsPriceRequest Request Object
type ListResourceSpecsPriceRequest struct {

	// 语言，非必填，忽略大小写，默认zh_cn，枚举：zh_cn：中文 en_us：英文
	XLanguage *string `json:"X-Language,omitempty"`

	Body *ResourceSpecsPriceQueryReq `json:"body,omitempty"`
}

func (o ListResourceSpecsPriceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourceSpecsPriceRequest struct{}"
	}

	return strings.Join([]string{"ListResourceSpecsPriceRequest", string(data)}, " ")
}
