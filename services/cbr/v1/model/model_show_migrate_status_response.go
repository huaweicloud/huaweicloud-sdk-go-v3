package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowMigrateStatusResponse Response Object
type ShowMigrateStatusResponse struct {

	// 租户迁移状态
	Status *ShowMigrateStatusResponseStatus `json:"status,omitempty"`

	// 项目迁移状态
	ProjectStatus  *[]DomainMigrateProjectStatus `json:"project_status,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o ShowMigrateStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMigrateStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowMigrateStatusResponse", string(data)}, " ")
}

type ShowMigrateStatusResponseStatus struct {
	value string
}

type ShowMigrateStatusResponseStatusEnum struct {
	SUCCESS   ShowMigrateStatusResponseStatus
	FAILED    ShowMigrateStatusResponseStatus
	MIGRATING ShowMigrateStatusResponseStatus
}

func GetShowMigrateStatusResponseStatusEnum() ShowMigrateStatusResponseStatusEnum {
	return ShowMigrateStatusResponseStatusEnum{
		SUCCESS: ShowMigrateStatusResponseStatus{
			value: "success",
		},
		FAILED: ShowMigrateStatusResponseStatus{
			value: "failed",
		},
		MIGRATING: ShowMigrateStatusResponseStatus{
			value: "migrating",
		},
	}
}

func (c ShowMigrateStatusResponseStatus) Value() string {
	return c.value
}

func (c ShowMigrateStatusResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowMigrateStatusResponseStatus) UnmarshalJSON(b []byte) error {
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
