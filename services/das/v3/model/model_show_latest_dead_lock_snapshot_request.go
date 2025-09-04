package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowLatestDeadLockSnapshotRequest Request Object
type ShowLatestDeadLockSnapshotRequest struct {

	// 连接id
	ConnectionId string `json:"connection_id"`

	// 死锁id
	Id int32 `json:"id"`

	// 语言
	XLanguage *ShowLatestDeadLockSnapshotRequestXLanguage `json:"X-Language,omitempty"`
}

func (o ShowLatestDeadLockSnapshotRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLatestDeadLockSnapshotRequest struct{}"
	}

	return strings.Join([]string{"ShowLatestDeadLockSnapshotRequest", string(data)}, " ")
}

type ShowLatestDeadLockSnapshotRequestXLanguage struct {
	value string
}

type ShowLatestDeadLockSnapshotRequestXLanguageEnum struct {
	ZH_CN ShowLatestDeadLockSnapshotRequestXLanguage
	EN_US ShowLatestDeadLockSnapshotRequestXLanguage
}

func GetShowLatestDeadLockSnapshotRequestXLanguageEnum() ShowLatestDeadLockSnapshotRequestXLanguageEnum {
	return ShowLatestDeadLockSnapshotRequestXLanguageEnum{
		ZH_CN: ShowLatestDeadLockSnapshotRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ShowLatestDeadLockSnapshotRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ShowLatestDeadLockSnapshotRequestXLanguage) Value() string {
	return c.value
}

func (c ShowLatestDeadLockSnapshotRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowLatestDeadLockSnapshotRequestXLanguage) UnmarshalJSON(b []byte) error {
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
