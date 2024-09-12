package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SharedConfig 模板共享配置。该配置仅用于模板。
type SharedConfig struct {

	// **参数解释**： 共享类型。 **约束限制**： 该配置仅用于模板 **取值范围**： * PRIVATE：私有，仅本租户可访问。 * PUBLIC：公开，所有租户可访问。当前仅提供系统资产可公开访问。 * SHARED：共享，指定租户可访问。拥有者指定租户可访问。  **默认取值**： 不涉及。
	SharedType *SharedConfigSharedType `json:"shared_type,omitempty"`

	// **参数解释**： 共享状态。 **约束限制**： 该配置仅用于shared_type为SHARED的模板。 **取值范围**： * PUBLISHED：发布。模板可用。 * DRAFT：草稿。编辑态，仅拥有者可访问。 * REVIEW：审核态。不可编辑，仅拥有者/审核人员可查看。  **默认取值**： 不涉及。
	SharedState *SharedConfigSharedState `json:"shared_state,omitempty"`

	// **参数解释**： 共享过期时间。空表示永久不过期。 **约束限制**： 该配置仅用于shared_type为SHARED的模板。 格式遵循：RFC 3339，示例“2021*01*10T08:43:17Z”。 **取值范围**： 字符长度0-20位 **默认取值**： 不涉及。
	ExpireTime *string `json:"expire_time,omitempty"`

	// **参数解释**： 允许访问本资产的租户列表。 **约束限制**： 该配置仅用于shared_type为SHARED的模板。 **取值范围**： 最大支持100个租户，重复的记录会被忽略。 租户ID填写project_id，字符长度1-64位。 **默认取值**： 不涉及。
	AllowedProjectIds *[]string `json:"allowed_project_ids,omitempty"`
}

func (o SharedConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SharedConfig struct{}"
	}

	return strings.Join([]string{"SharedConfig", string(data)}, " ")
}

type SharedConfigSharedType struct {
	value string
}

type SharedConfigSharedTypeEnum struct {
	PRIVATE SharedConfigSharedType
	PUBLIC  SharedConfigSharedType
	SHARED  SharedConfigSharedType
}

func GetSharedConfigSharedTypeEnum() SharedConfigSharedTypeEnum {
	return SharedConfigSharedTypeEnum{
		PRIVATE: SharedConfigSharedType{
			value: "PRIVATE",
		},
		PUBLIC: SharedConfigSharedType{
			value: "PUBLIC",
		},
		SHARED: SharedConfigSharedType{
			value: "SHARED",
		},
	}
}

func (c SharedConfigSharedType) Value() string {
	return c.value
}

func (c SharedConfigSharedType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SharedConfigSharedType) UnmarshalJSON(b []byte) error {
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

type SharedConfigSharedState struct {
	value string
}

type SharedConfigSharedStateEnum struct {
	PUBLISHED SharedConfigSharedState
	DRAFT     SharedConfigSharedState
	REVIEW    SharedConfigSharedState
}

func GetSharedConfigSharedStateEnum() SharedConfigSharedStateEnum {
	return SharedConfigSharedStateEnum{
		PUBLISHED: SharedConfigSharedState{
			value: "PUBLISHED",
		},
		DRAFT: SharedConfigSharedState{
			value: "DRAFT",
		},
		REVIEW: SharedConfigSharedState{
			value: "REVIEW",
		},
	}
}

func (c SharedConfigSharedState) Value() string {
	return c.value
}

func (c SharedConfigSharedState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SharedConfigSharedState) UnmarshalJSON(b []byte) error {
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
