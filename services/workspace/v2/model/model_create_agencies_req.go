package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateAgenciesReq 开通委托功能请求体。
type CreateAgenciesReq struct {

	// 委托场景。   - WORKSPACE：云桌面。   - CLOUD_GAME：云游戏。   - CLOUD_STORAGE：云存储。   - SCREEN_RECORD：录屏审计。
	Scene *CreateAgenciesReqScene `json:"scene,omitempty"`

	// 操作类型。 - CREATE 创建 - FIX 修复
	Action *CreateAgenciesReqAction `json:"action,omitempty"`
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
	WORKSPACE     CreateAgenciesReqScene
	CLOUD_GAME    CreateAgenciesReqScene
	CLOUD_STORAGE CreateAgenciesReqScene
	SCREEN_RECORD CreateAgenciesReqScene
}

func GetCreateAgenciesReqSceneEnum() CreateAgenciesReqSceneEnum {
	return CreateAgenciesReqSceneEnum{
		WORKSPACE: CreateAgenciesReqScene{
			value: "WORKSPACE",
		},
		CLOUD_GAME: CreateAgenciesReqScene{
			value: "CLOUD_GAME",
		},
		CLOUD_STORAGE: CreateAgenciesReqScene{
			value: "CLOUD_STORAGE",
		},
		SCREEN_RECORD: CreateAgenciesReqScene{
			value: "SCREEN_RECORD",
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

type CreateAgenciesReqAction struct {
	value string
}

type CreateAgenciesReqActionEnum struct {
	CREATE CreateAgenciesReqAction
	FIX    CreateAgenciesReqAction
}

func GetCreateAgenciesReqActionEnum() CreateAgenciesReqActionEnum {
	return CreateAgenciesReqActionEnum{
		CREATE: CreateAgenciesReqAction{
			value: "CREATE",
		},
		FIX: CreateAgenciesReqAction{
			value: "FIX",
		},
	}
}

func (c CreateAgenciesReqAction) Value() string {
	return c.value
}

func (c CreateAgenciesReqAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateAgenciesReqAction) UnmarshalJSON(b []byte) error {
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
