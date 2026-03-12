package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListWdrSnapshotsCollectResultsRequest Request Object
type ListWdrSnapshotsCollectResultsRequest struct {

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**:   - zh-cn   - en-us  **默认取值**: en-us
	XLanguage *ListWdrSnapshotsCollectResultsRequestXLanguage `json:"X-Language,omitempty"`

	// **参数解释**: 实例ID，此参数是用户创建实例的唯一标识。 **约束限制**: 不涉及。 **取值范围**: 只能由英文字母、数字组成，且长度为36个字符。 **默认取值**: 不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**: 索引位置，偏移量。 **约束限制**:   不涉及。 **取值范围**:   必须为数字，不能为负数。 **默认取值**:   默认为0。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**:   查询记录数。， **约束限制**:   不涉及。 **取值范围**:   不能为负数，最小值为1，最大值为100。 **默认取值**:   默认为10。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**:   查询开始时间。格式为“yyyy-mm-ddThh:mm:ssZ”。其中，T指某个时间的开始，Z指时区偏移量，时区中的+号需要进行URL编码，编码为%2B，时区中的-号无需编码。   例如北京时间偏移显示为+0800，start_time=2024-03-15T17:20:33+0800传参时编码为start_time=2024-03-15T17:20:33%2B0800。 **约束限制**:   不涉及。 **取值范围**:   不涉及。 **默认取值**:   不涉及。
	StartTime *string `json:"start_time,omitempty"`

	// **参数解释**:   查询结束时间。格式为“yyyy-mm-ddThh:mm:ssZ”。其中，T指某个时间的开始，Z指时区偏移量，时区中的+号需要进行URL编码，编码为%2B，时区中的-号无需编码。   例如北京时间偏移显示为+0800，end_time=2024-03-16T17:20:33+0800传参时编码为end_time=2024-03-16T17:20:33%2B0800。 **约束限制**:   不涉及。 **取值范围**:   不涉及。 **默认取值**:   不涉及。
	EndTime *string `json:"end_time,omitempty"`

	// **参数解释**:   任务ID。正确填写后，可查询指定任务对应的快照报告采集结果。不支持模糊匹配，需填写完整的任务ID。 **约束限制**:   不涉及。 **取值范围**:   不涉及。 **默认取值**:   不涉及。
	JobId *string `json:"job_id,omitempty"`

	// **参数解释**:   采集任务创建时间终点。可查询任务创建时间小于等于该时间终点的任务结果。格式为“yyyy-mm-ddThh:mm:ssZ”。其中，T指某个时间的开始；Z指时区偏移量。时区中的+号需要进行URL编码，编码为%2B，时区中的-号无需编码。例如，北京时间偏移显示为+0800，job_end_time=2024-03-16T17:20:33+0800，传参时编码为job_end_time=2024-03-16T17:20:33%2B0800。 **约束限制**:   不涉及。 **取值范围**:   不涉及。 **默认取值**:   不涉及。
	JobEndTime *string `json:"job_end_time,omitempty"`

	// **参数解释**:   采集任务创建时间起点。可查询任务创建时间大于等于该时间起点的任务结果。格式为“yyyy-mm-ddThh:mm:ssZ”。其中，T指某个时间的开始；Z指时区偏移量。时区中的+号需要进行URL编码，编码为%2B，时区中的-号无需编码。例如，北京时间偏移显示为+0800，job_start_time=2024-03-15T17:20:33+0800传参时编码为job_start_time=2024-03-15T17:20:33%2B0800。 **约束限制**:   不涉及。 **取值范围**:   不涉及。 **默认取值**:   不涉及。
	JobStartTime *string `json:"job_start_time,omitempty"`

	// **参数解释**: 任务采集状态。填写后，可查询对应采集状态的任务结果。 **约束限制**: 不支持模糊匹配，区分大小写，请填写完整的合法状态值。 **取值范围**: - EXPORTING：采集中。 - SUCCESS：采集成功。 - FAILED：采集失败。  **默认取值**:   不涉及。
	Status *string `json:"status,omitempty"`

	// **参数解释**: 填写后，可查询对应采集类型的任务结果。 **约束限制**: 不支持模糊匹配，区分大小写，请填写完整的合法枚举值。 **取值范围**: - cluster：实例级。 - component：组件级。 - pdb：租户级。  **默认取值**:   不涉及。
	WdrType *string `json:"wdr_type,omitempty"`
}

func (o ListWdrSnapshotsCollectResultsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWdrSnapshotsCollectResultsRequest struct{}"
	}

	return strings.Join([]string{"ListWdrSnapshotsCollectResultsRequest", string(data)}, " ")
}

type ListWdrSnapshotsCollectResultsRequestXLanguage struct {
	value string
}

type ListWdrSnapshotsCollectResultsRequestXLanguageEnum struct {
	ZH_CN ListWdrSnapshotsCollectResultsRequestXLanguage
	EN_US ListWdrSnapshotsCollectResultsRequestXLanguage
}

func GetListWdrSnapshotsCollectResultsRequestXLanguageEnum() ListWdrSnapshotsCollectResultsRequestXLanguageEnum {
	return ListWdrSnapshotsCollectResultsRequestXLanguageEnum{
		ZH_CN: ListWdrSnapshotsCollectResultsRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListWdrSnapshotsCollectResultsRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListWdrSnapshotsCollectResultsRequestXLanguage) Value() string {
	return c.value
}

func (c ListWdrSnapshotsCollectResultsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListWdrSnapshotsCollectResultsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
