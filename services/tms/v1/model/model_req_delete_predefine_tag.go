package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// 删除预定义标签请求
type ReqDeletePredefineTag struct {
	// 操作标识（区分大小写）：delete（删除）

	Action ReqDeletePredefineTagAction `json:"action"`
	// 标签列表

	Tags []PredefineTagRequest `json:"tags"`
}

func (o ReqDeletePredefineTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReqDeletePredefineTag struct{}"
	}

	return strings.Join([]string{"ReqDeletePredefineTag", string(data)}, " ")
}

type ReqDeletePredefineTagAction struct {
	value string
}

type ReqDeletePredefineTagActionEnum struct {
	DELETE ReqDeletePredefineTagAction
}

func GetReqDeletePredefineTagActionEnum() ReqDeletePredefineTagActionEnum {
	return ReqDeletePredefineTagActionEnum{
		DELETE: ReqDeletePredefineTagAction{
			value: "delete",
		},
	}
}

func (c ReqDeletePredefineTagAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ReqDeletePredefineTagAction) UnmarshalJSON(b []byte) error {
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
