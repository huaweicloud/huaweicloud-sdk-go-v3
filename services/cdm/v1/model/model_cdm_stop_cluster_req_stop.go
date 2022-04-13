package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// 集群启动操作，定义集群启动标识，为空对象
type CdmStopClusterReqStop struct {
	// 关机类型： - IMMEDIATELY：立即关机。 - GRACEFULLY：优雅关机。

	StopMode CdmStopClusterReqStopStopMode `json:"stopMode"`
	// 关机时延，仅在stopMode为“GRACEFULLY”生效，单位：秒。该值为-1时，表示等待所有作业完成，并停止接受新作业。该值为大于0的任意值表示等待该时长后关机，并停止接受新作业。

	DelayTime *int32 `json:"delayTime,omitempty"`
}

func (o CdmStopClusterReqStop) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CdmStopClusterReqStop struct{}"
	}

	return strings.Join([]string{"CdmStopClusterReqStop", string(data)}, " ")
}

type CdmStopClusterReqStopStopMode struct {
	value string
}

type CdmStopClusterReqStopStopModeEnum struct {
	IMMEDIATELY CdmStopClusterReqStopStopMode
	GRACEFULLY  CdmStopClusterReqStopStopMode
}

func GetCdmStopClusterReqStopStopModeEnum() CdmStopClusterReqStopStopModeEnum {
	return CdmStopClusterReqStopStopModeEnum{
		IMMEDIATELY: CdmStopClusterReqStopStopMode{
			value: "IMMEDIATELY",
		},
		GRACEFULLY: CdmStopClusterReqStopStopMode{
			value: "GRACEFULLY",
		},
	}
}

func (c CdmStopClusterReqStopStopMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CdmStopClusterReqStopStopMode) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
