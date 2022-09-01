package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type UpdateIaConfigResponse struct {

	// 配置ID
	Id *string `json:"id,omitempty" xml:"id"`

	// 配置项名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 配置项详情
	Value *string `json:"value,omitempty" xml:"value"`

	// 配置项描述
	Description *string `json:"description,omitempty" xml:"description"`

	// 版本号
	Version *int64 `json:"version,omitempty" xml:"version"`

	// 下发状态
	State *UpdateIaConfigResponseState `json:"state,omitempty" xml:"state"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty" xml:"create_time"`

	// 更新时间
	UpdateTime     *string `json:"update_time,omitempty" xml:"update_time"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateIaConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateIaConfigResponse struct{}"
	}

	return strings.Join([]string{"UpdateIaConfigResponse", string(data)}, " ")
}

type UpdateIaConfigResponseState struct {
	value string
}

type UpdateIaConfigResponseStateEnum struct {
	SUCCESS UpdateIaConfigResponseState
	SENDING UpdateIaConfigResponseState
}

func GetUpdateIaConfigResponseStateEnum() UpdateIaConfigResponseStateEnum {
	return UpdateIaConfigResponseStateEnum{
		SUCCESS: UpdateIaConfigResponseState{
			value: "SUCCESS",
		},
		SENDING: UpdateIaConfigResponseState{
			value: "SENDING",
		},
	}
}

func (c UpdateIaConfigResponseState) Value() string {
	return c.value
}

func (c UpdateIaConfigResponseState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateIaConfigResponseState) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
