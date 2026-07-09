package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResourceSpecsRequest Request Object
type ListResourceSpecsRequest struct {

	// |参数名称：语言| |参数的约束及描述：非必填，忽略大小写，默认zh_cn，枚举：zh_cn：中文 en_us：英文|
	XLanguage *string `json:"X-Language,omitempty"`

	Body *ResourceSpecsQueryReq `json:"body,omitempty"`
}

func (o ListResourceSpecsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourceSpecsRequest struct{}"
	}

	return strings.Join([]string{"ListResourceSpecsRequest", string(data)}, " ")
}
