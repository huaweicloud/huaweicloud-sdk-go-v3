package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BackboneBandwidthResp 骨干带宽对象
type BackboneBandwidthResp struct {

	// 骨干带宽的uuid
	Id *string `json:"id,omitempty"`

	// 骨干带宽的状态
	AdminStatus *string `json:"admin_status,omitempty"`

	// 骨干带宽的大小
	Size *int32 `json:"size,omitempty"`

	// 骨干带宽的short_id
	ShortId *string `json:"short_id,omitempty"`

	// 描述网络等级，从高到低分为铂金、金、银、铜
	SlaLevel *BackboneBandwidthRespSlaLevel `json:"sla_level,omitempty"`

	// 线路质量金银铜对应的DSCP值
	Dscp *int32 `json:"dscp,omitempty"`
}

func (o BackboneBandwidthResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackboneBandwidthResp struct{}"
	}

	return strings.Join([]string{"BackboneBandwidthResp", string(data)}, " ")
}

type BackboneBandwidthRespSlaLevel struct {
	value string
}

type BackboneBandwidthRespSlaLevelEnum struct {
	PT BackboneBandwidthRespSlaLevel
	AU BackboneBandwidthRespSlaLevel
	AG BackboneBandwidthRespSlaLevel
	CU BackboneBandwidthRespSlaLevel
}

func GetBackboneBandwidthRespSlaLevelEnum() BackboneBandwidthRespSlaLevelEnum {
	return BackboneBandwidthRespSlaLevelEnum{
		PT: BackboneBandwidthRespSlaLevel{
			value: "Pt",
		},
		AU: BackboneBandwidthRespSlaLevel{
			value: "Au",
		},
		AG: BackboneBandwidthRespSlaLevel{
			value: "Ag",
		},
		CU: BackboneBandwidthRespSlaLevel{
			value: "Cu",
		},
	}
}

func (c BackboneBandwidthRespSlaLevel) Value() string {
	return c.value
}

func (c BackboneBandwidthRespSlaLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BackboneBandwidthRespSlaLevel) UnmarshalJSON(b []byte) error {
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
