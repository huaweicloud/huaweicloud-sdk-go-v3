package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBackupRestoreTimeRequest Request Object
type ShowBackupRestoreTimeRequest struct {

	// 语言。
	XLanguage *string `json:"X-Language,omitempty"`

	// 租户在某一project下的实例ID。
	InstanceId string `json:"instance_id"`

	// **参数解释**：  所需查询的日期。  **约束限制**：  不涉及。  **取值范围**：  yyyy-mm-dd字符串格式，时区为UTC。  **默认取值**：  不涉及。
	Date *string `json:"date,omitempty"`

	// **参数解释**：  所需查询的起始时间戳。  **约束限制**：  - “start_time”有值时，“end_time”必选。 - “date”有值时，“start_time”失效。  **取值范围**：  格式为UNIX时间戳，单位是毫秒，时区为UTC标准时区。传参时需要将对应时区的时间转为标准时区对应的时间戳，比如，北京时区的时间点需要-8h后再转为时间戳。  **默认取值**：  不涉及。
	StartTime *string `json:"start_time,omitempty"`

	// **参数解释**：  所需查询的结束时间戳。  **约束限制**：  - “end_time”有值时，“start_time”必选。 - “date”有值时，“end_time”失效。  **取值范围**：  格式为UNIX时间戳，单位是毫秒，时区为UTC标准时区。传参时需要将对应时区的时间转为标准时区对应的时间戳，比如，北京时区的时间点需要-8h后再转为时间戳。  **默认取值**：  不涉及。
	EndTime *string `json:"end_time,omitempty"`
}

func (o ShowBackupRestoreTimeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBackupRestoreTimeRequest struct{}"
	}

	return strings.Join([]string{"ShowBackupRestoreTimeRequest", string(data)}, " ")
}
