package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Region 支持的Region。
type Region struct {

	// 实例名称。
	Name string `json:"name"`

	AreaId *AreaIdDef `json:"area_id"`

	// Region ID。
	Id *string `json:"id,omitempty"`

	// 大区名。
	AreaName *string `json:"area_name,omitempty"`

	// 云连接使用场景 er vpc vgw。
	UsedScenes *[]RegionUsedScenes `json:"used_scenes,omitempty"`
}

func (o Region) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Region struct{}"
	}

	return strings.Join([]string{"Region", string(data)}, " ")
}

type RegionUsedScenes struct {
	value string
}

type RegionUsedScenesEnum struct {
	ER  RegionUsedScenes
	VPC RegionUsedScenes
	VGW RegionUsedScenes
}

func GetRegionUsedScenesEnum() RegionUsedScenesEnum {
	return RegionUsedScenesEnum{
		ER: RegionUsedScenes{
			value: "er",
		},
		VPC: RegionUsedScenes{
			value: "vpc",
		},
		VGW: RegionUsedScenes{
			value: "vgw",
		},
	}
}

func (c RegionUsedScenes) Value() string {
	return c.value
}

func (c RegionUsedScenes) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RegionUsedScenes) UnmarshalJSON(b []byte) error {
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
