package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ApiVersionInfo This is a auto create Response Object
type ApiVersionInfo struct {

	// API版本号。  取值：由高到低版本分别为v3，v2，v2.0。
	Id string `json:"id"`

	// API版本的状态。  取值： - CURRENT：当前版本。 - STABLE：稳定版本。 - DEPRECATED：废弃版本。  说明： 所有支持的API版本中最高版状态为CURRENT，其他版本状态为STABLE。
	Status ApiVersionInfoStatus `json:"status"`
}

func (o ApiVersionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiVersionInfo struct{}"
	}

	return strings.Join([]string{"ApiVersionInfo", string(data)}, " ")
}

type ApiVersionInfoStatus struct {
	value string
}

type ApiVersionInfoStatusEnum struct {
	CURRENT    ApiVersionInfoStatus
	STABLE     ApiVersionInfoStatus
	DEPRECATED ApiVersionInfoStatus
}

func GetApiVersionInfoStatusEnum() ApiVersionInfoStatusEnum {
	return ApiVersionInfoStatusEnum{
		CURRENT: ApiVersionInfoStatus{
			value: "CURRENT",
		},
		STABLE: ApiVersionInfoStatus{
			value: "STABLE",
		},
		DEPRECATED: ApiVersionInfoStatus{
			value: "DEPRECATED",
		},
	}
}

func (c ApiVersionInfoStatus) Value() string {
	return c.value
}

func (c ApiVersionInfoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiVersionInfoStatus) UnmarshalJSON(b []byte) error {
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
