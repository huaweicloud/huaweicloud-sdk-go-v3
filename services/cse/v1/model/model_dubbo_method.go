package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DubboMethod struct {

	// 服务方法。
	ServiceMethod *string `json:"serviceMethod,omitempty"`

	// 附加请求头。
	HeadersAttach *string `json:"headersAttach,omitempty"`

	// http 方法。
	HttpMethods *[]string `json:"httpMethods,omitempty"`

	// http 路径。
	HttpPath *string `json:"httpPath,omitempty"`

	// dubbo 方法参数。
	Params *[]DubboMethodParam `json:"params,omitempty"`
}

func (o DubboMethod) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DubboMethod struct{}"
	}

	return strings.Join([]string{"DubboMethod", string(data)}, " ")
}
