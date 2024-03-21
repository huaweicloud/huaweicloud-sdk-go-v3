package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowPlaybookMonitorsRequest Request Object
type ShowPlaybookMonitorsRequest struct {

	// application/json;charset=UTF-8
	ContentType string `json:"content-type"`

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 剧本ID
	PlaybookId string `json:"playbook_id"`

	// 开始时间。格式ISO8601：YYYY-MM-DDTHH:mm:ss.ms+timezone。例如：2021-01-30T23:00:00Z+0800。时区信息为剧本实例产生的时区，无法解析时区的时间，默认时区填东八区。
	StartTime string `json:"start_time"`

	// 统计剧本版本类型（ALL:全部，VALID:有效的，DELETED:已删除）
	VersionQueryType ShowPlaybookMonitorsRequestVersionQueryType `json:"version_query_type"`

	// 结束时间。格式ISO8601：YYYY-MM-DDTHH:mm:ss.ms+timezone。例如：2021-01-30T23:00:00Z+0800。时区信息为剧本实例产生的时区，无法解析时区的时间，默认时区填东八区。
	EndTime string `json:"end_time"`
}

func (o ShowPlaybookMonitorsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPlaybookMonitorsRequest struct{}"
	}

	return strings.Join([]string{"ShowPlaybookMonitorsRequest", string(data)}, " ")
}

type ShowPlaybookMonitorsRequestVersionQueryType struct {
	value string
}

type ShowPlaybookMonitorsRequestVersionQueryTypeEnum struct {
	ALLVALIDDELETED ShowPlaybookMonitorsRequestVersionQueryType
}

func GetShowPlaybookMonitorsRequestVersionQueryTypeEnum() ShowPlaybookMonitorsRequestVersionQueryTypeEnum {
	return ShowPlaybookMonitorsRequestVersionQueryTypeEnum{
		ALLVALIDDELETED: ShowPlaybookMonitorsRequestVersionQueryType{
			value: "ALL:全部，VALID:有效的，DELETED:已删除",
		},
	}
}

func (c ShowPlaybookMonitorsRequestVersionQueryType) Value() string {
	return c.value
}

func (c ShowPlaybookMonitorsRequestVersionQueryType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowPlaybookMonitorsRequestVersionQueryType) UnmarshalJSON(b []byte) error {
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
