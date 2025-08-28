package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAgenciesV5Request Request Object
type ListAgenciesV5Request struct {

	// 每页显示的条目数量，范围为1到200条，默认为100条。
	Limit *int32 `json:"limit,omitempty"`

	// 分页标记，长度为4到400个字符，只包含字母、数字、\"+\"、\"/\"、\"=\"、\"-\"和\"_\"的字符串。
	Marker *string `json:"marker,omitempty"`

	// 资源路径前缀，由若干段字符串拼接而成，每段先包含一个或多个字母、数字、\".\"、\",\"、\"+\"、\"@\"、\"=\"、\"_\"或\"-\"，并以\"/\"结尾，例如\"foo/bar/\"。
	PathPrefix *string `json:"path_prefix,omitempty"`
}

func (o ListAgenciesV5Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAgenciesV5Request struct{}"
	}

	return strings.Join([]string{"ListAgenciesV5Request", string(data)}, " ")
}
