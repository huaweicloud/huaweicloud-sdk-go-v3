package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListProtectionServersRequest Request Object
type ListProtectionServersRequest struct {

	// 企业项目ID
	EnterpriseProjectId string `json:"enterprise_project_id"`

	// 查询起始点
	Offset int32 `json:"offset"`

	// 每页显示个数
	Limit int32 `json:"limit"`

	// 云主机名称
	HostName *string `json:"host_name,omitempty"`

	// 操作系统类型，包含如下2种。  - linux ：linux类型应用防护。  - windows ：windows类型应用防护。
	OsType *string `json:"os_type,omitempty"`

	// 云主机私有IP
	HostIp *string `json:"host_ip,omitempty"`

	// 应用类型，包含如下1种。   - java ：java类型应用防护。
	AppType *ListProtectionServersRequestAppType `json:"app_type,omitempty"`

	// 应用防护状态，包含如下2种。   - closed ：未开启。   - opened ：防护中。
	AppStatus ListProtectionServersRequestAppStatus `json:"app_status"`
}

func (o ListProtectionServersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProtectionServersRequest struct{}"
	}

	return strings.Join([]string{"ListProtectionServersRequest", string(data)}, " ")
}

type ListProtectionServersRequestAppType struct {
	value string
}

type ListProtectionServersRequestAppTypeEnum struct {
	JAVA ListProtectionServersRequestAppType
}

func GetListProtectionServersRequestAppTypeEnum() ListProtectionServersRequestAppTypeEnum {
	return ListProtectionServersRequestAppTypeEnum{
		JAVA: ListProtectionServersRequestAppType{
			value: "java",
		},
	}
}

func (c ListProtectionServersRequestAppType) Value() string {
	return c.value
}

func (c ListProtectionServersRequestAppType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListProtectionServersRequestAppType) UnmarshalJSON(b []byte) error {
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

type ListProtectionServersRequestAppStatus struct {
	value string
}

type ListProtectionServersRequestAppStatusEnum struct {
	CLOSED ListProtectionServersRequestAppStatus
	OPENED ListProtectionServersRequestAppStatus
}

func GetListProtectionServersRequestAppStatusEnum() ListProtectionServersRequestAppStatusEnum {
	return ListProtectionServersRequestAppStatusEnum{
		CLOSED: ListProtectionServersRequestAppStatus{
			value: "closed",
		},
		OPENED: ListProtectionServersRequestAppStatus{
			value: "opened",
		},
	}
}

func (c ListProtectionServersRequestAppStatus) Value() string {
	return c.value
}

func (c ListProtectionServersRequestAppStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListProtectionServersRequestAppStatus) UnmarshalJSON(b []byte) error {
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
