package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type NodeconfigStatus struct {

	// **参数解释**： 插件实例的状态。 **取值范围**： 可选值如下： - Pending：安装中，表示插件正在安装中。 - Running：运行中，表示插件全部实例状态都在运行中，插件正常使用。 - Updating：升级中，表示插件正在更新中。 - Abnormal：不可用，表示插件状态异常，插件不可使用。可单击状态查看失败原因。 - Deleting：删除中，表示插件正在删除中。
	Phase NodeconfigStatusPhase `json:"phase"`
}

func (o NodeconfigStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeconfigStatus struct{}"
	}

	return strings.Join([]string{"NodeconfigStatus", string(data)}, " ")
}

type NodeconfigStatusPhase struct {
	value string
}

type NodeconfigStatusPhaseEnum struct {
	PENDING  NodeconfigStatusPhase
	UPDATING NodeconfigStatusPhase
	RUNNING  NodeconfigStatusPhase
	ABNORMAL NodeconfigStatusPhase
	DELETING NodeconfigStatusPhase
}

func GetNodeconfigStatusPhaseEnum() NodeconfigStatusPhaseEnum {
	return NodeconfigStatusPhaseEnum{
		PENDING: NodeconfigStatusPhase{
			value: "Pending",
		},
		UPDATING: NodeconfigStatusPhase{
			value: "Updating",
		},
		RUNNING: NodeconfigStatusPhase{
			value: "Running",
		},
		ABNORMAL: NodeconfigStatusPhase{
			value: "Abnormal",
		},
		DELETING: NodeconfigStatusPhase{
			value: "Deleting",
		},
	}
}

func (c NodeconfigStatusPhase) Value() string {
	return c.value
}

func (c NodeconfigStatusPhase) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NodeconfigStatusPhase) UnmarshalJSON(b []byte) error {
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
