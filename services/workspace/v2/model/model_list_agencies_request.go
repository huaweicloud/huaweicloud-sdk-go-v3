package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListAgenciesRequest Request Object
type ListAgenciesRequest struct {

	// 委托场景。 - WORKSPACE：云桌面。 - CLOUD_GAME：云游戏。 - SCREEN_RECORD：录屏审计。
	Scene *ListAgenciesRequestScene `json:"scene,omitempty"`
}

func (o ListAgenciesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAgenciesRequest struct{}"
	}

	return strings.Join([]string{"ListAgenciesRequest", string(data)}, " ")
}

type ListAgenciesRequestScene struct {
	value string
}

type ListAgenciesRequestSceneEnum struct {
	WORKSPACE     ListAgenciesRequestScene
	CLOUD_GAME    ListAgenciesRequestScene
	SCREEN_RECORD ListAgenciesRequestScene
}

func GetListAgenciesRequestSceneEnum() ListAgenciesRequestSceneEnum {
	return ListAgenciesRequestSceneEnum{
		WORKSPACE: ListAgenciesRequestScene{
			value: "WORKSPACE",
		},
		CLOUD_GAME: ListAgenciesRequestScene{
			value: "CLOUD_GAME",
		},
		SCREEN_RECORD: ListAgenciesRequestScene{
			value: "SCREEN_RECORD",
		},
	}
}

func (c ListAgenciesRequestScene) Value() string {
	return c.value
}

func (c ListAgenciesRequestScene) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAgenciesRequestScene) UnmarshalJSON(b []byte) error {
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
