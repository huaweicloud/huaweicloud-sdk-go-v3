package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CollectTranscriberJobResponse struct {

	// 当前识别状态。具体状态如下所示：  WAITING 等待识别。 FINISHED 识别已经完成。 ERROR 识别过程中发生错误。
	Status *string `json:"status,omitempty" xml:"status"`

	// 任务创建时间, 遵循 RFC 3339格式。 格式示例：2018-12-04T13:10:29.310Z。
	CreateTime *string `json:"create_time,omitempty" xml:"create_time"`

	// 开始识别时间, 遵循 RFC 3339格式。 当status为FINISHED或ERROR时存在。 格式示例：2018-12-04T13:10:29.310Z。
	StartTime *string `json:"start_time,omitempty" xml:"start_time"`

	// 识别完成时间, 遵循 RFC 3339格式。 当status为FINISHED或ERROR时存在。 格式示例：2018-12-04T13:10:29.310Z。
	FinishTime *string `json:"finish_time,omitempty" xml:"finish_time"`

	// 转写结果, 多句结果的数组。
	Segments *[]Segment `json:"segments,omitempty" xml:"segments"`

	// 音频时长，单位ms
	AudioDuration  *int32 `json:"audio_duration,omitempty" xml:"audio_duration"`
	HttpStatusCode int    `json:"-"`
}

func (o CollectTranscriberJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CollectTranscriberJobResponse struct{}"
	}

	return strings.Join([]string{"CollectTranscriberJobResponse", string(data)}, " ")
}
