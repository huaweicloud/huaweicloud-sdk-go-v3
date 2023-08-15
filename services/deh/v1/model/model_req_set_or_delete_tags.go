package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ReqSetOrDeleteTags 批量添加的标签。
type ReqSetOrDeleteTags struct {

	// 标签列表。
	Tags []ResourceTag `json:"tags"`

	// 操作标识（仅支持小写）：create（创建），delete（删除）。
	Action ReqSetOrDeleteTagsAction `json:"action"`
}

func (o ReqSetOrDeleteTags) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReqSetOrDeleteTags struct{}"
	}

	return strings.Join([]string{"ReqSetOrDeleteTags", string(data)}, " ")
}

type ReqSetOrDeleteTagsAction struct {
	value string
}

type ReqSetOrDeleteTagsActionEnum struct {
	CREATE ReqSetOrDeleteTagsAction
	DELETE ReqSetOrDeleteTagsAction
}

func GetReqSetOrDeleteTagsActionEnum() ReqSetOrDeleteTagsActionEnum {
	return ReqSetOrDeleteTagsActionEnum{
		CREATE: ReqSetOrDeleteTagsAction{
			value: "create",
		},
		DELETE: ReqSetOrDeleteTagsAction{
			value: "delete",
		},
	}
}

func (c ReqSetOrDeleteTagsAction) Value() string {
	return c.value
}

func (c ReqSetOrDeleteTagsAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ReqSetOrDeleteTagsAction) UnmarshalJSON(b []byte) error {
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
