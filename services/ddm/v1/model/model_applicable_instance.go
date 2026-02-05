package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApplicableInstance struct {

	// **参数解释**：  实例ID，此参数是实例的唯一标识。  **参数范围**：  只能由英文字母、数字组成，后缀为in09，长度为36个字符。
	InstanceId *string `json:"instance_id,omitempty"`

	// **参数解释**：  实例名称。  **参数范围**：  不涉及。
	InstanceName *string `json:"instance_name,omitempty"`
}

func (o ApplicableInstance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplicableInstance struct{}"
	}

	return strings.Join([]string{"ApplicableInstance", string(data)}, " ")
}
