package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPrivateProviderVersionMetadataRequest Request Object
type ShowPrivateProviderVersionMetadataRequest struct {

	// 用户指定的，对于此请求的唯一Id，用于定位某个请求，推荐使用UUID
	ClientRequestId string `json:"Client-Request-Id"`

	// 私有provider（private-provider）的名称。此名字在domain_id+region下应唯一，可以使用小写英文、数字、中划线。仅支持以小写英文、数字开头结尾。
	ProviderName string `json:"provider_name"`

	// provider的版本号。版本号遵循语义化版本号（Semantic Version），为用户自定义
	ProviderVersion string `json:"provider_version"`

	// 私有provider的唯一Id，由资源编排服务随机生成，为UUID。  由于私有provider名称仅仅在同一时间下唯一，即用户允许先生成一个叫helloword的私有provider，删除，再重新创建一个同名私有provider。  对于团队并行开发，用户可能希望确保，当前我操作的私有provider就是我以为的那个，而不是由其他队友删除后创建的同名私有provider。  因此，使用ID就可以做到强匹配。资源编排服务保证每次创建私有provider所对应的Id都不相同。  如果给予的provider_id和当前私有provider的Id不一致，则返回400。
	ProviderId *string `json:"provider_id,omitempty"`
}

func (o ShowPrivateProviderVersionMetadataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPrivateProviderVersionMetadataRequest struct{}"
	}

	return strings.Join([]string{"ShowPrivateProviderVersionMetadataRequest", string(data)}, " ")
}
