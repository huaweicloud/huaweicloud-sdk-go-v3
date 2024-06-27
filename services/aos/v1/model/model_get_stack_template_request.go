package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetStackTemplateRequest Request Object
type GetStackTemplateRequest struct {

	// 用户指定的，对于此请求的唯一ID，用于定位某个请求，推荐使用UUID
	ClientRequestId string `json:"Client-Request-Id"`

	// 资源栈的名称。此名字在domain_id+区域+project_id下应唯一，可以使用中文、大小写英文、数字、下划线、中划线。首字符需为中文或者英文，区分大小写。
	StackName string `json:"stack_name"`

	// 资源栈（stack）的唯一Id。  此Id由资源编排服务在生成资源栈的时候生成，为UUID。  由于资源栈名仅仅在同一时间下唯一，即用户允许先生成一个叫HelloWorld的资源栈，删除，再重新创建一个同名资源栈。  对于团队并行开发，用户可能希望确保，当前我操作的资源栈就是我认为的那个，而不是其他队友删除后创建的同名资源栈。因此，使用ID就可以做到强匹配。  资源编排服务保证每次创建的资源栈所对应的ID都不相同，更新不会影响ID。如果给予的stack_id和当前资源栈的ID不一致，则返回400
	StackId *string `json:"stack_id,omitempty"`

	// 允许访问资源栈模板的source ip列表，source ip应是具有CIDR表示法且带有子网掩码的IPv4地址。
	AccessControlSourceIps *[]string `json:"access_control_source_ips,omitempty"`

	// 允许访问资源栈模板的source vpc id列表， source vpc id应仅包含小写字母、数字或中划线。
	AccessControlSourceVpcIds *[]string `json:"access_control_source_vpc_ids,omitempty"`
}

func (o GetStackTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetStackTemplateRequest struct{}"
	}

	return strings.Join([]string{"GetStackTemplateRequest", string(data)}, " ")
}
