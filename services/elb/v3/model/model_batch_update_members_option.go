package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateMembersOption 批量添加member请求参数。
type BatchUpdateMembersOption struct {

	// 后端服务器ID。 >此处并非ECS服务器的ID，而是ELB为绑定的后端服务器自动生成的member ID。
	Id string `json:"id"`

	// 后端云服务器的管理状态。取值：true、false。  虽然创建、更新请求支持该字段，但实际取值决定于后端云服务器对应的弹性云服务器是否存在。若存在，该值为true，否则，该值为false。  请勿传入该字段。
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	// 后端服务器名称。
	Name *string `json:"name,omitempty"`

	// 后端服务器端口。  在开启端口透传的pool下的member，该字段无法更新。  [网关型LB，即pool协议为IP时，protocol_port必须设置为0。](tag:hws_eu)
	ProtocolPort *int32 `json:"protocol_port,omitempty"`

	// 后端云服务器的权重，请求按权重在同一后端云服务器组下的后端云服务器间分发。权重为0的后端不再接受新的请求。当后端云服务器所在的后端云服务器组的lb_algorithm的取值为SOURCE_IP时，该字段无效。
	Weight *int32 `json:"weight,omitempty"`
}

func (o BatchUpdateMembersOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateMembersOption struct{}"
	}

	return strings.Join([]string{"BatchUpdateMembersOption", string(data)}, " ")
}
