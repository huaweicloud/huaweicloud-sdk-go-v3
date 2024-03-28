package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSlDomainSettingV2Request Request Object
type UpdateSlDomainSettingV2Request struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
	InstanceId string `json:"instance_id"`

	// 分组的编号
	GroupId string `json:"group_id"`

	Body *SlDomainAccessSetting `json:"body,omitempty"`
}

func (o UpdateSlDomainSettingV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSlDomainSettingV2Request struct{}"
	}

	return strings.Join([]string{"UpdateSlDomainSettingV2Request", string(data)}, " ")
}
