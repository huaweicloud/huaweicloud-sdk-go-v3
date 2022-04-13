package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 更新后端云服务器的请求体
type UpdateMemberReq struct {
	// 后端云服务器名称。

	Name *string `json:"name,omitempty"`
	// 后端云服务器的管理状态；该字段虽然支持创建、更新，但实际取值决定于后端云服务器对应的弹性云服务器是否存在。该字段虽然支持创建、更新，但实际取值决定于member对应的弹性云服务器是否存在。若存在，该值为true，否则，该值为false。

	AdminStateUp *bool `json:"admin_state_up,omitempty"`
	// 后端云服务器的权重，请求按权重在同一后端云服务器组下的后端云服务器间分发。权重为0的后端不再接受新的请求。当后端云服务器所在的后端云服务器组的lb_algorithm的取值为SOURCE_IP时，该字段无效。

	Weight *int32 `json:"weight,omitempty"`
}

func (o UpdateMemberReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMemberReq struct{}"
	}

	return strings.Join([]string{"UpdateMemberReq", string(data)}, " ")
}
