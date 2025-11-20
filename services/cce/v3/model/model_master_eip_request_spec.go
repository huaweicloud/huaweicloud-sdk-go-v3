package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type MasterEipRequestSpec struct {

	// **参数解释**： 绑定或解绑动作 **约束限制**： 不涉及 **取值范围**： - bind：绑定 - unbind：解绑 **默认取值**： 不涉及
	Action *MasterEipRequestSpecAction `json:"action,omitempty"`

	Spec *MasterEipRequestSpecSpec `json:"spec,omitempty"`

	// **参数解释**： 带宽(字段已失效，暂不推荐使用) **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Bandwidth *string `json:"bandwidth,omitempty"`

	// **参数解释**： 弹性网卡IP(字段已失效，暂不推荐使用) **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	ElasticIp *string `json:"elasticIp,omitempty"`
}

func (o MasterEipRequestSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MasterEipRequestSpec struct{}"
	}

	return strings.Join([]string{"MasterEipRequestSpec", string(data)}, " ")
}

type MasterEipRequestSpecAction struct {
	value string
}

type MasterEipRequestSpecActionEnum struct {
	BIND   MasterEipRequestSpecAction
	UNBIND MasterEipRequestSpecAction
}

func GetMasterEipRequestSpecActionEnum() MasterEipRequestSpecActionEnum {
	return MasterEipRequestSpecActionEnum{
		BIND: MasterEipRequestSpecAction{
			value: "bind",
		},
		UNBIND: MasterEipRequestSpecAction{
			value: "unbind",
		},
	}
}

func (c MasterEipRequestSpecAction) Value() string {
	return c.value
}

func (c MasterEipRequestSpecAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MasterEipRequestSpecAction) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
