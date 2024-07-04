package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddCustomIngressPortRequest Request Object
type AddCustomIngressPortRequest struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
	InstanceId string `json:"instance_id"`

	Body *IngressPortCreate `json:"body,omitempty"`
}

func (o AddCustomIngressPortRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddCustomIngressPortRequest struct{}"
	}

	return strings.Join([]string{"AddCustomIngressPortRequest", string(data)}, " ")
}
