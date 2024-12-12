package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PrivateModuleVersion struct {

	// 私有模块（private-module）的名字。此名字在domain_id+region下应唯一，可以使用中文、大小写英文、数字、下划线、中划线。首字符需为中文或者英文，区分大小写。
	ModuleName string `json:"module_name"`

	// 私有模块（private-module）的唯一Id。  此Id由资源编排服务在生成模块的时候生成，为UUID。  由于私有模块名仅仅在同一时间下唯一，即用户允许先生成一个叫HelloWorld的私有模块，删除，在重新创建一个同名私有模块。  对于团队并行开发，用户可能希望确保，当前我操作的私有模块就是我以为的那个，而不是其他队友删除后创建的同名私有模块。因此，使用Id就可以做到强匹配。  资源编排服务保证每次创建的私有模块所对应的Id都不相同，更新不会影响Id。如果给予的module_id和当前模块的Id不一致，则返回400
	ModuleId *string `json:"module_id,omitempty"`

	// 模块的版本号。版本号必须遵循语义化版本号（Semantic Version），为用户自定义
	ModuleVersion *string `json:"module_version,omitempty"`

	// 模块版本（module version）的描述。可用于客户识别并管理模块的版本。注意：模块版本为不可更新（immutable），即描述不可更新，如果需要更新，请删除后重建
	VersionDescription *string `json:"version_description,omitempty"`

	// 私有模块（private-module）版本的生成时间，格式为YYYY-MM-DDTHH:mm:ss.SSSZ，精确到毫秒，UTC时区，即，如1970-01-01T00:00:00.000Z。
	CreateTime *string `json:"create_time,omitempty"`

	// 在模板中使用模块需要定义如下格式：   module \"my_hello_word_module\" {     source = {module_source}   }  其中{module_source}为本参数
	ModuleSource *string `json:"module_source,omitempty"`
}

func (o PrivateModuleVersion) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrivateModuleVersion struct{}"
	}

	return strings.Join([]string{"PrivateModuleVersion", string(data)}, " ")
}
