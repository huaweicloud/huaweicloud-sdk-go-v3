package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePrivateProviderResponse Response Object
type CreatePrivateProviderResponse struct {

	// 私有provider（private-provider）的唯一Id。  此Id由资源编排服务在生成provider的时候生成，为UUID。  由于私有provider名仅仅在同一时间下唯一，即用户允许先生成一个叫HelloWorld的私有provider，删除，再重新创建一个同名私有provider。  对于团队并行开发，用户可能希望确保，当前我操作的私有provider就是我以为的那个，而不是其他队友删除后创建的同名私有provider。因此，使用Id就可以做到强匹配。  资源编排服务保证每次创建的私有provider所对应的Id都不相同，更新不会影响Id。如果给予的provider_id和当前provider的Id不一致，则返回400
	ProviderId *string `json:"provider_id,omitempty"`

	// 用户使用私有provider，在Terraform模板中定义required_providers信息时，需要指明的source参数。  该参数按照“huawei.com/private-provider/{provider_name}”的形式拼接。关于provider_name和provider_source字段在模板中的使用细节，详见创建私有Provider的API描述。
	ProviderSource *string `json:"provider_source,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreatePrivateProviderResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePrivateProviderResponse struct{}"
	}

	return strings.Join([]string{"CreatePrivateProviderResponse", string(data)}, " ")
}
