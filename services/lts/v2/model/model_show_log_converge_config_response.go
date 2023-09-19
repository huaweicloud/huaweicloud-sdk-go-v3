package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowLogConvergeConfigResponse Response Object
type ShowLogConvergeConfigResponse struct {

	// ID
	Id *string `json:"id,omitempty"`

	// 组织成员账号id
	MemberAccountId *string `json:"member_account_id,omitempty"`

	// 管理员或者委托管理员项目id
	MemberProjectId *string `json:"member_project_id,omitempty"`

	// 创建时间
	CreateTime *int64 `json:"create_time,omitempty"`

	// 更新时间
	UpdateTime *int64 `json:"update_time,omitempty"`

	// creating: 配置创建中   done：配置创建完成
	Status *ShowLogConvergeConfigResponseStatus `json:"status,omitempty"`

	// 组织id
	OrganizationId *string `json:"organization_id,omitempty"`

	// 管理员或者委托管理员账号id
	ManagementAccountId *string `json:"management_account_id,omitempty"`

	// 管理员项目id
	ManagementProjectId *string `json:"management_project_id,omitempty"`

	// 版本
	Version *string `json:"version,omitempty"`

	// 日志汇聚配置
	LogMappingConfig *[]LogMappingConfig `json:"log_mapping_config,omitempty"`
	HttpStatusCode   int                 `json:"-"`
}

func (o ShowLogConvergeConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLogConvergeConfigResponse struct{}"
	}

	return strings.Join([]string{"ShowLogConvergeConfigResponse", string(data)}, " ")
}

type ShowLogConvergeConfigResponseStatus struct {
	value string
}

type ShowLogConvergeConfigResponseStatusEnum struct {
	CREATING ShowLogConvergeConfigResponseStatus
	DONE     ShowLogConvergeConfigResponseStatus
}

func GetShowLogConvergeConfigResponseStatusEnum() ShowLogConvergeConfigResponseStatusEnum {
	return ShowLogConvergeConfigResponseStatusEnum{
		CREATING: ShowLogConvergeConfigResponseStatus{
			value: "creating",
		},
		DONE: ShowLogConvergeConfigResponseStatus{
			value: " done",
		},
	}
}

func (c ShowLogConvergeConfigResponseStatus) Value() string {
	return c.value
}

func (c ShowLogConvergeConfigResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowLogConvergeConfigResponseStatus) UnmarshalJSON(b []byte) error {
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
