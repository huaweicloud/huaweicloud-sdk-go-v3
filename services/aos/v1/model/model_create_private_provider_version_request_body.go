package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreatePrivateProviderVersionRequestBody struct {

	// 私有provider（private-provider）的唯一Id。  此Id由资源编排服务在生成provider的时候生成，为UUID。  由于私有provider名仅仅在同一时间下唯一，即用户允许先生成一个叫HelloWorld的私有provider，删除，再重新创建一个同名私有provider。  对于团队并行开发，用户可能希望确保，当前我操作的私有provider就是我以为的那个，而不是其他队友删除后创建的同名私有provider。因此，使用Id就可以做到强匹配。  资源编排服务保证每次创建的私有provider所对应的Id都不相同，更新不会影响Id。如果给予的provider_id和当前provider的Id不一致，则返回400
	ProviderId *string `json:"provider_id,omitempty"`

	// provider的版本号。版本号必须遵循语义化版本号（Semantic Version），为用户自定义
	ProviderVersion string `json:"provider_version"`

	// 私有provider版本（provider version）的描述。可用于客户识别并管理私有provider的版本。注意：provider版本为不可更新（immutable），所以该字段不可更新，如果需要更新，请删除后重建
	VersionDescription *string `json:"version_description,omitempty"`

	// FunctionGraph方法的统一资源标识，用于标识唯一的FunctionGraph方法。当前只支持和RFS同region的function_graph_urn，如果给予了关于其他region的，会报错400。  关于该参数的详细解释，请参考官方文档：https://support.huaweicloud.com/api-functiongraph/functiongraph_06_0102.html
	FunctionGraphUrn string `json:"function_graph_urn"`
}

func (o CreatePrivateProviderVersionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePrivateProviderVersionRequestBody struct{}"
	}

	return strings.Join([]string{"CreatePrivateProviderVersionRequestBody", string(data)}, " ")
}
