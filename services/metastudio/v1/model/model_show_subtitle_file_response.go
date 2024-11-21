package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSubtitleFileResponse Response Object
type ShowSubtitleFileResponse struct {

	// 字幕文件所有任务状态。 * GENERATING：字幕文件任务生成中。 * GENERATED：字幕文件生成结束。
	JobState *string `json:"job_state,omitempty"`

	// 字幕文件列表。
	SubtitleFileInfo *[]SubtitleFileDetail `json:"subtitle_file_info,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowSubtitleFileResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSubtitleFileResponse struct{}"
	}

	return strings.Join([]string{"ShowSubtitleFileResponse", string(data)}, " ")
}
