package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateAgenciesReq 开通委托功能请求体。
type CreateAgenciesReq struct {

	// 委托场景。 - WORKSPACE：云桌面。 - CLOUD_GAME：云游戏。
	Scene *CreateAgenciesReqScene `json:"scene,omitempty"`
}

func (o CreateAgenciesReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAgenciesReq struct{}"
	}

	return strings.Join([]string{"CreateAgenciesReq", string(data)}, " ")
}

type CreateAgenciesReqScene struct {
	value string
}

type CreateAgenciesReqSceneEnum struct {
	WORKSPACE  CreateAgenciesReqScene
	CLOUD_GAME CreateAgenciesReqScene
}

func GetCreateAgenciesReqSceneEnum() CreateAgenciesReqSceneEnum {
	return CreateAgenciesReqSceneEnum{
		WORKSPACE: CreateAgenciesReqScene{
			value: "WORKSPACE",
		},
		CLOUD_GAME: CreateAgenciesReqScene{
			value: "CLOUD_GAME",
		},
	}
}

func (c CreateAgenciesReqScene) Value() string {
	return c.value
}

func (c CreateAgenciesReqScene) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateAgenciesReqScene) UnmarshalJSON(b []byte) error {
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
