package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExtraInfo struct {

	// 子级导入的包
	ChildImportPackage *[]string `json:"childImportPackage,omitempty"`

	// 文档链接
	DocumentLink *string `json:"document_link,omitempty"`

	// http请求方法
	HttpMethod *string `json:"http_method,omitempty"`

	// HTTP请求的URL
	HttpUrl *string `json:"http_url,omitempty"`

	// 导入的包
	ImportPackage *[]string `json:"importPackage,omitempty"`

	Mock *MockInfo `json:"mock,omitempty"`

	// 输出参数
	OutputParam *[]AwVariable `json:"outputParam,omitempty"`

	// 参数依赖
	ParamDependent *string `json:"param_dependent,omitempty"`

	// 摘要
	Summary *string `json:"summary,omitempty"`
}

func (o ExtraInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExtraInfo struct{}"
	}

	return strings.Join([]string{"ExtraInfo", string(data)}, " ")
}
