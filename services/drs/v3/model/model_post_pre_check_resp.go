package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

//
type PostPreCheckResp struct {
	// 任务id。

	Id *string `json:"id,omitempty"`
	// 预检查id。

	PrecheckId *string `json:"precheck_id,omitempty"`
	// 成功或失败的状态

	Status *PostPreCheckRespStatus `json:"status,omitempty"`
	// 错误码。

	ErrorCode *string `json:"error_code,omitempty"`
	// 错误信息。

	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o PostPreCheckResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PostPreCheckResp struct{}"
	}

	return strings.Join([]string{"PostPreCheckResp", string(data)}, " ")
}

type PostPreCheckRespStatus struct {
	value string
}

type PostPreCheckRespStatusEnum struct {
	SUCCESS PostPreCheckRespStatus
	FAILED  PostPreCheckRespStatus
}

func GetPostPreCheckRespStatusEnum() PostPreCheckRespStatusEnum {
	return PostPreCheckRespStatusEnum{
		SUCCESS: PostPreCheckRespStatus{
			value: "success",
		},
		FAILED: PostPreCheckRespStatus{
			value: "failed",
		},
	}
}

func (c PostPreCheckRespStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PostPreCheckRespStatus) UnmarshalJSON(b []byte) error {
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
