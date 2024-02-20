package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdateFindingsReqBody struct {

	// 要更新的调查结果的ID。
	Ids *[]string `json:"ids,omitempty"`

	// 唯一的资源名称。
	ResourceUrn *string `json:"resource_urn,omitempty"`

	// 状态表示为更新查找状态而要采取的操作。使用“存档”将活动查找更改为存档查找。使用“活动”将存档的查找更改为活动查找。
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
