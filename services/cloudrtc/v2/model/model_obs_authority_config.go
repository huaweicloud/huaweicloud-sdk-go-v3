package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ObsAuthorityConfig struct {

	// OBS桶名
	Bucket string `json:"bucket"`

	// 操作，1-授权；0-取消授权
	Operation int32 `json:"operation"`

	// 查询bucket所在的region - cn-north-4 - cn-north-1 - cn-north-5 - cn-north-6 - cn-south-1 - cn-east-2
	Location ObsAuthorityConfigLocation `json:"location"`

	// 租户华为云账号projectid
	ProjectId *string `json:"project_id,omitempty"`
}

func (o ObsAuthorityConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ObsAuthorityConfig struct{}"
	}

	return strings.Join([]string{"ObsAuthorityConfig", string(data)}, " ")
}

type ObsAuthorityConfigLocation struct {
	value string
}

type ObsAuthorityConfigLocationEnum struct {
	CN_NORTH_4 ObsAuthorityConfigLocation
	CN_NORTH_1 ObsAuthorityConfigLocation
	CN_NORTH_5 ObsAuthorityConfigLocation
	CN_NORTH_6 ObsAuthorityConfigLocation
	CN_SOUTH_1 ObsAuthorityConfigLocation
	CN_EAST_2  ObsAuthorityConfigLocation
}

func GetObsAuthorityConfigLocationEnum() ObsAuthorityConfigLocationEnum {
	return ObsAuthorityConfigLocationEnum{
		CN_NORTH_4: ObsAuthorityConfigLocation{
			value: "cn-north-4",
		},
		CN_NORTH_1: ObsAuthorityConfigLocation{
			value: "cn-north-1",
		},
		CN_NORTH_5: ObsAuthorityConfigLocation{
			value: "cn-north-5",
		},
		CN_NORTH_6: ObsAuthorityConfigLocation{
			value: "cn-north-6",
		},
		CN_SOUTH_1: ObsAuthorityConfigLocation{
			value: "cn-south-1",
		},
		CN_EAST_2: ObsAuthorityConfigLocation{
			value: "cn-east-2",
		},
	}
}

func (c ObsAuthorityConfigLocation) Value() string {
	return c.value
}

func (c ObsAuthorityConfigLocation) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ObsAuthorityConfigLocation) UnmarshalJSON(b []byte) error {
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
