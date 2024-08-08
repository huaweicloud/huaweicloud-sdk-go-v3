package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeInstanceChargeModeRequest Request Object
type ChangeInstanceChargeModeRequest struct {

	// **参数说明**：实例ID。 **取值范围**：长度不超过36，由小写字母[a-f]、数字、连接符（-）的组成。
	InstanceId string `json:"instance_id"`

	Body *ChangeInstanceChargeMode `json:"body,omitempty"`
}

func (o ChangeInstanceChargeModeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeInstanceChargeModeRequest struct{}"
	}

	return strings.Join([]string{"ChangeInstanceChargeModeRequest", string(data)}, " ")
}
