package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBusinessDiscountInfoRequest Request Object
type ListBusinessDiscountInfoRequest struct {

	// 忽略大小写，中文：zh_cn 英文：en_us。缺省为zh_cn
	XLanguage *string `json:"X-Language,omitempty"`

	Body *BusinessDiscountQueryReq `json:"body,omitempty"`
}

func (o ListBusinessDiscountInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBusinessDiscountInfoRequest struct{}"
	}

	return strings.Join([]string{"ListBusinessDiscountInfoRequest", string(data)}, " ")
}
