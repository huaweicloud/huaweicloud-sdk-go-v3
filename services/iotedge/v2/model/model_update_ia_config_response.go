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
	Id *string `json:"id,omitempty"`

	// 配置项名称
	Name *string `json:"name,omitempty"`

	// 配置项详情
	Value *string `json:"value,omitempty"`

	// 配置项描述
	Description *string `json:"description,omitempty"`

	// 版本号
	Version *int64 `json:"version,omitempty"`

	// 下发状态
	State *UpdateIaConfigResponseState `json:"state,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 更新时间
	UpdateTime     *string `json:"update_time,omitempty"`
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
