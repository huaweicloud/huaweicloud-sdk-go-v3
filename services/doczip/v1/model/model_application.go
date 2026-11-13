package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

type Application struct {

	// **参数解释**： 应用编码。 **取值范围**： [1-256]字符。
	Id *string `json:"id,omitempty"`

	// **参数解释**： 应用名称。 **取值范围**： [1-200]字符。
	Name string `json:"name"`

	// **参数解释**： 应用图标URL，当前仅支持预置应用图标，取值为0~9的数字，分别标识第1到第10张预置图标。 **取值范围**： [1-200]字符。
	Icon *string `json:"icon,omitempty"`

	// **参数解释**： 应用描述。 **取值范围**： [1-200]字符。
	Description *string `json:"description,omitempty"`

	// **参数解释**： 发布说明。 **取值范围**： [1-500]字符。
	ReleaseNote *string `json:"release_note,omitempty"`

	// **参数解释**： [创建时间。格式为YYYY-MM-DD HH:mm:ss。例如：2023-11-24 17:35:25。](tag:his) [创建时间。UTC时间。格式为YYYY-MM-DD HH:mm:ss。例如：2023-11-24 17:35:25。](tag:hws) **取值范围**： 不涉及。
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	// **参数解释**： [更新时间。格式为YYYY-MM-DD HH:mm:ss。例如：2023-11-24 17:35:25。](tag:his) [更新时间。UTC时间。格式为YYYY-MM-DD HH:mm:ss。例如：2023-11-24 17:35:25。](tag:hws) **取值范围**： 不涉及。
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`

	// **参数解释**： 创建用户ID。 **取值范围**： [1-50]个字符。
	CreateUserId *string `json:"create_user_id,omitempty"`

	// **参数解释**： 更新用户ID。 **取值范围**： [1-50]个字符。
	UpdateUserId *string `json:"update_user_id,omitempty"`

	// **参数解释**： 应用状态。 **取值范围**： - 0：待发布。 - 1：已发布，生效状态。 - 2：已发布，失效状态。
	ReleaseState *ApplicationReleaseState `json:"release_state,omitempty"`

	// **参数解释**： 版本号。 **取值范围**： [1-50]个字符。
	Version *string `json:"version,omitempty"`
}

func (o Application) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Application struct{}"
	}

	return strings.Join([]string{"Application", string(data)}, " ")
}

type ApplicationReleaseState struct {
	value int32
}

type ApplicationReleaseStateEnum struct {
	E_0 ApplicationReleaseState
	E_1 ApplicationReleaseState
	E_2 ApplicationReleaseState
}

func GetApplicationReleaseStateEnum() ApplicationReleaseStateEnum {
	return ApplicationReleaseStateEnum{
		E_0: ApplicationReleaseState{
			value: 0,
		}, E_1: ApplicationReleaseState{
			value: 1,
		}, E_2: ApplicationReleaseState{
			value: 2,
		},
	}
}

func (c ApplicationReleaseState) Value() int32 {
	return c.value
}

func (c ApplicationReleaseState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApplicationReleaseState) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
