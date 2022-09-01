package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type BatchDeleteClusterTagsReq struct {

	// 操作标识：仅限于delete（删除）。
	Action BatchDeleteClusterTagsReqAction `json:"action" xml:"action"`

	// 标签列表。
	Tags []Tag `json:"tags" xml:"tags"`
}

func (o BatchDeleteClusterTagsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteClusterTagsReq struct{}"
	}

	return strings.Join([]string{"BatchDeleteClusterTagsReq", string(data)}, " ")
}

type BatchDeleteClusterTagsReqAction struct {
	value string
}

type BatchDeleteClusterTagsReqActionEnum struct {
	DELETE BatchDeleteClusterTagsReqAction
}

func GetBatchDeleteClusterTagsReqActionEnum() BatchDeleteClusterTagsReqActionEnum {
	return BatchDeleteClusterTagsReqActionEnum{
		DELETE: BatchDeleteClusterTagsReqAction{
			value: "delete",
		},
	}
}

func (c BatchDeleteClusterTagsReqAction) Value() string {
	return c.value
}

func (c BatchDeleteClusterTagsReqAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchDeleteClusterTagsReqAction) UnmarshalJSON(b []byte) error {
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
