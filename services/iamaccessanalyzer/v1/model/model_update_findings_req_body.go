package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdateFindingsReqBody struct {

	// 要更新的访问分析结果唯一标识符数组。
	Ids *[]string `json:"ids,omitempty"`

	// 资源的唯一资源标识符。
	ResourceUrn *string `json:"resource_urn,omitempty"`

	// 要更新的访问分析结果状态。
	Status UpdateFindingsReqBodyStatus `json:"status"`
}

func (o UpdateFindingsReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFindingsReqBody struct{}"
	}

	return strings.Join([]string{"UpdateFindingsReqBody", string(data)}, " ")
}

type UpdateFindingsReqBodyStatus struct {
	value string
}

type UpdateFindingsReqBodyStatusEnum struct {
	ACTIVE   UpdateFindingsReqBodyStatus
	ARCHIVED UpdateFindingsReqBodyStatus
}

func GetUpdateFindingsReqBodyStatusEnum() UpdateFindingsReqBodyStatusEnum {
	return UpdateFindingsReqBodyStatusEnum{
		ACTIVE: UpdateFindingsReqBodyStatus{
			value: "active",
		},
		ARCHIVED: UpdateFindingsReqBodyStatus{
			value: "archived",
		},
	}
}

func (c UpdateFindingsReqBodyStatus) Value() string {
	return c.value
}

func (c UpdateFindingsReqBodyStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateFindingsReqBodyStatus) UnmarshalJSON(b []byte) error {
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
