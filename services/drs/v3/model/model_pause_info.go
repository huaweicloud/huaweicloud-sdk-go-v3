package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 暂停任务请求参数
type PauseInfo struct {

	// 任务id
	JobId string `json:"job_id" xml:"job_id"`

	// 暂停类型，target:停回放,all:停日志抓取和回放
	PauseMode PauseInfoPauseMode `json:"pause_mode" xml:"pause_mode"`
}

func (o PauseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PauseInfo struct{}"
	}

	return strings.Join([]string{"PauseInfo", string(data)}, " ")
}

type PauseInfoPauseMode struct {
	value string
}

type PauseInfoPauseModeEnum struct {
	TARGET PauseInfoPauseMode
	ALL    PauseInfoPauseMode
}

func GetPauseInfoPauseModeEnum() PauseInfoPauseModeEnum {
	return PauseInfoPauseModeEnum{
		TARGET: PauseInfoPauseMode{
			value: "target",
		},
		ALL: PauseInfoPauseMode{
			value: "all",
		},
	}
}

func (c PauseInfoPauseMode) Value() string {
	return c.value
}

func (c PauseInfoPauseMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PauseInfoPauseMode) UnmarshalJSON(b []byte) error {
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
