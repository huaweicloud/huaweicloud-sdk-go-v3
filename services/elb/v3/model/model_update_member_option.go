package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 更新后端服务器请求参数。
type UpdateMemberOption struct {
	// 后端云服务器的管理状态。取值：true、false。  虽然创建、更新请求支持该字段，但实际取值决定于后端云服务器对应的弹性云服务器是否存在。若存在，该值为true，否则，该值为false。  请勿传入该字段。

	AdminStateUp *bool `json:"admin_state_up,omitempty"`
	// 后端云服务器名称。

	Name *string `json:"name,omitempty"`
	// 后端云服务器的权重，请求将根据pool配置的负载均衡算法和后端云服务器的权重进行负载分发。权重值越大，分发的请求越多。权重为0的后端不再接受新的请求。 取值：0-100，默认1。 使用说明：  - 若所在pool的lb_algorithm取值为SOURCE_IP，该字段无效。

	Weight *int32 `json:"weight,omitempty"`
}

func (o UpdateMemberOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMemberOption struct{}"
	}

	return strings.Join([]string{"UpdateMemberOption", string(data)}, " ")
}
