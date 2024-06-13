package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CommitJobReq 用户提交任务请求
type CommitJobReq struct {
	Tag *JobTag `json:"tag,omitempty"`

	// 一段描述信息,会呈现在资产库中。
	Description *string `json:"description,omitempty"`

	// 语音性别,是男性声音还是女性声音。 * FEMALE: 女性 * MALE: 男性
	Sex *CommitJobReqSex `json:"sex,omitempty"`

	// 音色名称。该名称会作为资产库中音色模型资产名称。
	VoiceName *string `json:"voice_name,omitempty"`

	// 训练语言,当前仅支持中文。 * CN: 中文 * EN: 英文
	Language *string `json:"language,omitempty"`

	// 手机号
	Phone *string `json:"phone,omitempty"`

	// 第三方用户id
	AppUserId *string `json:"app_user_id,omitempty"`
}

func (o CommitJobReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommitJobReq struct{}"
	}

	return strings.Join([]string{"CommitJobReq", string(data)}, " ")
}

type CommitJobReqSex struct {
	value string
}

type CommitJobReqSexEnum struct {
	FEMALE CommitJobReqSex
	MALE   CommitJobReqSex
}

func GetCommitJobReqSexEnum() CommitJobReqSexEnum {
	return CommitJobReqSexEnum{
		FEMALE: CommitJobReqSex{
			value: "FEMALE",
		},
		MALE: CommitJobReqSex{
			value: "MALE",
		},
	}
}

func (c CommitJobReqSex) Value() string {
	return c.value
}

func (c CommitJobReqSex) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CommitJobReqSex) UnmarshalJSON(b []byte) error {
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
