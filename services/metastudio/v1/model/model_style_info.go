package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// StyleInfo 服务开通请求
type StyleInfo struct {

	// 数字人风格化名称
	Name string `json:"name"`

	// 数字人风格化描述
	Description *string `json:"description,omitempty"`

	// 租户ID
	ProjectId *string `json:"project_id,omitempty"`

	// 状态
	Status *string `json:"status,omitempty"`

	// 性别
	Sex *StyleInfoSex `json:"sex,omitempty"`

	// 标签。单个标签16字节，多个用逗号分隔，最多50个。
	Tags *[]string `json:"tags,omitempty"`

	// 风格化素材资产组合。
	StyleAssets *[]StyleAssetItem `json:"style_assets,omitempty"`

	ExtraMeta *StyleExtraMeta `json:"extra_meta,omitempty"`

	// 数字人风格ID
	StyleId *string `json:"style_id,omitempty"`

	// 数字人风格创建时间，格式遵循：RFC 3339。 例 “2020-07-30T10:43:17Z”。
	CreateTime *string `json:"create_time,omitempty"`

	// 数字人风格更新时间，格式遵循：RFC 3339。 例 “2020-07-30T10:43:17Z”。
	UpdateTime *string `json:"update_time,omitempty"`

	// 数字人风格状态枚举 * CREATING：创建中 * PUBLISHED：已发布 * DELETED：已删除 * UNPUBLISHED：未发布 * PUBLISHING：发布中
	State *StyleInfoState `json:"state,omitempty"`
}

func (o StyleInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StyleInfo struct{}"
	}

	return strings.Join([]string{"StyleInfo", string(data)}, " ")
}

type StyleInfoSex struct {
	value string
}

type StyleInfoSexEnum struct {
	UNKNOW StyleInfoSex
	MALE   StyleInfoSex
	FEMALE StyleInfoSex
}

func GetStyleInfoSexEnum() StyleInfoSexEnum {
	return StyleInfoSexEnum{
		UNKNOW: StyleInfoSex{
			value: "UNKNOW",
		},
		MALE: StyleInfoSex{
			value: "MALE",
		},
		FEMALE: StyleInfoSex{
			value: "FEMALE",
		},
	}
}

func (c StyleInfoSex) Value() string {
	return c.value
}

func (c StyleInfoSex) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StyleInfoSex) UnmarshalJSON(b []byte) error {
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

type StyleInfoState struct {
	value string
}

type StyleInfoStateEnum struct {
	CREATING    StyleInfoState
	PUBLISHED   StyleInfoState
	DELETED     StyleInfoState
	UNPUBLISHED StyleInfoState
	PUBLISHING  StyleInfoState
}

func GetStyleInfoStateEnum() StyleInfoStateEnum {
	return StyleInfoStateEnum{
		CREATING: StyleInfoState{
			value: "CREATING",
		},
		PUBLISHED: StyleInfoState{
			value: "PUBLISHED",
		},
		DELETED: StyleInfoState{
			value: "DELETED",
		},
		UNPUBLISHED: StyleInfoState{
			value: "UNPUBLISHED",
		},
		PUBLISHING: StyleInfoState{
			value: "PUBLISHING",
		},
	}
}

func (c StyleInfoState) Value() string {
	return c.value
}

func (c StyleInfoState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StyleInfoState) UnmarshalJSON(b []byte) error {
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
