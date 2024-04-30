package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreatePrivateProviderRequestBody struct {

	// 私有provider（private-provider）的名称。此名字在domain_id+region下应唯一，可以使用小写英文、数字、中划线。仅支持以小写英文、数字开头结尾。  按照HCL最佳实践，该名称推荐为在模板中定义的provider的本地名称（local_name）。  创建私有Provider（CreatePrivateProvider）API 还会以 “huawei.com/private-provider”为固定前缀，并以“huawei.com/private-provider/{provider_name}”的形式返回provider_source字段。关于provider_name和provider_source字段在模板中的使用细节，详见创建私有Provider的API描述。
	ProviderName string `json:"provider_name"`

	// 私有provider（private-provider）的描述。可用于客户识别被管理的私有provider。
	ProviderDescription *string `json:"provider_description,omitempty"`

	// provider的版本号。版本号必须遵循语义化版本号（Semantic Version），为用户自定义
	ProviderVersion *string `json:"provider_version,omitempty"`

	// 私有provider版本（provider version）的描述。可用于客户识别并管理私有provider的版本。注意：provider版本为不可更新（immutable），所以该字段不可更新，如果需要更新，请删除后重建
	VersionDescription *string `json:"version_description,omitempty"`

	// FunctionGraph方法的统一资源标识，用于唯一标识的FunctionGraph方法。当前只支持和RFS同region的function_graph_urn，如果给予了关于其他region的，会报错400。  关于该参数的详细解释，请参考官方文档：https://support.huaweicloud.com/api-functiongraph/functiongraph_06_0102.html
	FunctionGraphUrn *string `json:"function_graph_urn,omitempty"`
}

func (o CreatePrivateProviderRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePrivateProviderRequestBody struct{}"
	}

	return strings.Join([]string{"CreatePrivateProviderRequestBody", string(data)}, " ")
}
