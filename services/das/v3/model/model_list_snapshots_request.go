package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListSnapshotsRequest Request Object
type ListSnapshotsRequest struct {

	// 连接id
	ConnectionId string `json:"connection_id"`

	// 锁类型， 检查死锁为2
	Module int64 `json:"module"`

	// 开始时间
	StartAt int64 `json:"start_at"`

	// 结束时间
	EndAt int64 `json:"end_at"`

	// 每页返回的数量，默认值为10
	PerPage *int32 `json:"per_page,omitempty"`

	// 当前页码，第一次查询的时候默认值为1
	CurPage *int32 `json:"cur_page,omitempty"`

	// 语言
	XLanguage *ListSnapshotsRequestXLanguage `json:"X-Language,omitempty"`
}

func (o ListSnapshotsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSnapshotsRequest struct{}"
	}

	return strings.Join([]string{"ListSnapshotsRequest", string(data)}, " ")
}

type ListSnapshotsRequestXLanguage struct {
	value string
}

type ListSnapshotsRequestXLanguageEnum struct {
	ZH_CN ListSnapshotsRequestXLanguage
	EN_US ListSnapshotsRequestXLanguage
}

func GetListSnapshotsRequestXLanguageEnum() ListSnapshotsRequestXLanguageEnum {
	return ListSnapshotsRequestXLanguageEnum{
		ZH_CN: ListSnapshotsRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListSnapshotsRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListSnapshotsRequestXLanguage) Value() string {
	return c.value
}

func (c ListSnapshotsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSnapshotsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
