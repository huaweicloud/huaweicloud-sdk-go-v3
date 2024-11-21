package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PrivateHookSummary struct {

	// 私有hook的名字。此名字在domain_id+region下应唯一，可以使用中文、大小写英文、数字、下划线、中划线。首字符需为中文或者英文，区分大小写。  推荐用户使用三段命名空间：{自定义hook名称}-{hook应用场景}-hook。
	HookName string `json:"hook_name"`

	// 私有hook（private-hook）的唯一Id。  此Id由资源编排服务在生成私有hook的时候生成，为UUID。  由于私有hook名称仅仅在同一时间下唯一，即用户允许先生成一个叫HelloWorld的私有hook，删除，再重新创建一个同名私有hook。  对于团队并行开发，用户可能希望确保，当前我操作的私有hook就是我认为的那个，而不是其他队友删除后创建的同名私有hook。因此，使用Id就可以做到强匹配。  资源编排服务保证每次创建的私有hook所对应的Id都不相同，更新不会影响Id。如果给予的hook_id和当前hook的Id不一致，则返回400。
	HookId *string `json:"hook_id,omitempty"`

	// 私有hook的描述。可用于客户识别创建的私有hook。可通过UpdatePrivateHook API更新私有hook的描述。
	HookDescription *string `json:"hook_description,omitempty"`

	// 私有hook的默认版本号。版本号必须遵循语义化版本号（Semantic Version），为用户自定义。
	DefaultVersion *string `json:"default_version,omitempty"`

	Configuration *ConfigurationPrimitiveTypeHolderConfiguration `json:"configuration,omitempty"`

	// 私有Hook（private-hook）的生成时间，格式为YYYY-MM-DDTHH:mm:ss.SSSZ，精确到毫秒，UTC时区，即，如1970-01-01T00:00:00.000Z。
	CreateTime *string `json:"create_time,omitempty"`

	// 私有Hook（private-hook）的更新时间，格式为YYYY-MM-DDTHH:mm:ss.SSSZ，精确到毫秒，UTC时区，即，如1970-01-01T00:00:00.000Z。
	UpdateTime *string `json:"update_time,omitempty"`
}

func (o PrivateHookSummary) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrivateHookSummary struct{}"
	}

	return strings.Join([]string{"PrivateHookSummary", string(data)}, " ")
}
