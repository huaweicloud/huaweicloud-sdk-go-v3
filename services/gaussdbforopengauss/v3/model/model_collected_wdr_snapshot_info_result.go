package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CollectedWdrSnapshotInfoResult struct {

	// **参数解释**： 任务ID。 **取值范围**： 不涉及。
	JobId *string `json:"job_id,omitempty"`

	// **参数解释**： 采集类型。 **取值范围**： - 实例级则为cluster。 - 组件级则为component。
	WdrType *string `json:"wdr_type,omitempty"`

	// **参数解释**： 文件大小单位kb。当status为SUCCESS时，该值不为空。 **取值范围**： 不涉及。
	FileSize *int32 `json:"file_size,omitempty"`

	// **参数解释**： 下发采集时填写的开始snapshot时间。格式为“yyyy-mm-ddThh:mm:ssZ”。其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。 **取值范围**： 不涉及。
	StartTime *string `json:"start_time,omitempty"`

	// **参数解释**： 下发采集时填写的结束snapshot时间。格式为“yyyy-mm-ddThh:mm:ssZ”。其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。 **取值范围**： 不涉及。
	EndTime *string `json:"end_time,omitempty"`

	// **参数解释**： 报告下载链接，有效时间为30分钟。当status为SUCCESS时，该值不为空。 **取值范围**： 不涉及。
	DownloadUrl *string `json:"download_url,omitempty"`

	// **参数解释**： 采集状态。 **取值范围**: - SUCCESS：采集成功。 - FAILED：采集失败。 - EXPORTING：采集中。
	Status *string `json:"status,omitempty"`

	// **参数解释**： 备注。采集类型为组件级时，内容包括采集的组件ID。 **取值范围**： 不涉及。
	Notes *string `json:"notes,omitempty"`
}

func (o CollectedWdrSnapshotInfoResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CollectedWdrSnapshotInfoResult struct{}"
	}

	return strings.Join([]string{"CollectedWdrSnapshotInfoResult", string(data)}, " ")
}
