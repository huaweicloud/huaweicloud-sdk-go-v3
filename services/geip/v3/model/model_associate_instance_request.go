package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateInstanceRequest Request Object
type AssociateInstanceRequest struct {

	// 绑定接口可以加，标识请求是从哪个服务调过来的
	BindingInstanceService *string `json:"binding-instance-service,omitempty"`

	GlobalEipId string `json:"global_eip_id"`

	Body *AssociateInstanceGlobalEipRequestBody `json:"body,omitempty"`
}

func (o AssociateInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateInstanceRequest struct{}"
	}

	return strings.Join([]string{"AssociateInstanceRequest", string(data)}, " ")
}
