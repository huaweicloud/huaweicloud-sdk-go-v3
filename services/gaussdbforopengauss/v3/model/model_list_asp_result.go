package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListAspResult 采集的asp报告信息。
type ListAspResult struct {

	// **参数解释**: 任务ID。 **取值范围**: 不涉及。
	JobId string `json:"job_id"`

	// **参数解释**: 文件大小单位KB。 **取值范围**: 当status为SUCCESS时，该值不为空。
	FileSize *int32 `json:"file_size,omitempty"`

	// **参数解释**: 开始采集时间。 格式为“yyyy-mm-ddThh:mm:ssZ”。 其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。 **取值范围**: 不涉及。
	StartTime string `json:"start_time"`

	// **参数解释**: 结束采集时间。 格式为“yyyy-mm-ddThh:mm:ssZ”。 其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。 **取值范围**: 不涉及。
	EndTime string `json:"end_time"`

	// **参数解释**: 报告下载链接，有效时间为30分钟。 **取值范围**: 当status为SUCCESS时，该值不为空。
	DownloadUrl *string `json:"download_url,omitempty"`

	// **参数解释**: 采集状态。 **取值范围**: - SUCCESS：成功。 - FAILED：失败。 - EXPORTING：采集中。
	Status ListAspResultStatus `json:"status"`
}

func (o ListAspResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAspResult struct{}"
	}

	return strings.Join([]string{"ListAspResult", string(data)}, " ")
}

type ListAspResultStatus struct {
	value string
}

type ListAspResultStatusEnum struct {
	SUCCESS   ListAspResultStatus
	FAILED    ListAspResultStatus
	EXPORTING ListAspResultStatus
}

func GetListAspResultStatusEnum() ListAspResultStatusEnum {
	return ListAspResultStatusEnum{
		SUCCESS: ListAspResultStatus{
			value: "SUCCESS",
		},
		FAILED: ListAspResultStatus{
			value: "FAILED",
		},
		EXPORTING: ListAspResultStatus{
			value: "EXPORTING",
		},
	}
}

func (c ListAspResultStatus) Value() string {
	return c.value
}

func (c ListAspResultStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAspResultStatus) UnmarshalJSON(b []byte) error {
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
