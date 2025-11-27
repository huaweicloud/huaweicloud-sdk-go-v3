package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPrivateModuleVersionMetadataRequest Request Object
type ShowPrivateModuleVersionMetadataRequest struct {

	// 用户指定的，对于此请求的唯一Id，用于定位某个请求，推荐使用UUID
	ClientRequestId string `json:"Client-Request-Id"`

	// 私有模块（private-module）的名字。此名字在domain_id+region下应唯一，可以使用中文、大小写英文、数字、下划线、中划线。首字符需为中文或者英文，区分大小写。
	ModuleName string `json:"module_name"`

	// 私有模块（private-module）的唯一Id。  此Id由资源编排服务在生成模块的时候生成，为UUID。  由于私有模块名称仅仅在同一时间下唯一，即用户允许先生成一个叫HelloWorld的私有模块，删除，再重新创建一个同名私有模块。  对于团队并行开发，用户可能希望确保，当前我操作的私有模块就是我认为的那个，而不是其他队友删除后创建的同名私有模块。因此，使用Id就可以做到强匹配。  资源编排服务保证每次创建的私有模块所对应的Id都不相同，更新不会影响Id。如果给予的module_id和当前模块的Id不一致，则返回400
	ModuleId *string `json:"module_id,omitempty"`

	// 模块的版本号。版本号遵循语义化版本号（Semantic Version），为用户自定义
	ModuleVersion string `json:"module_version"`
}

func (o ShowPrivateModuleVersionMetadataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPrivateModuleVersionMetadataRequest struct{}"
	}

	return strings.Join([]string{"ShowPrivateModuleVersionMetadataRequest", string(data)}, " ")
}
