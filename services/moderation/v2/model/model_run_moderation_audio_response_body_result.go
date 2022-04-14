package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 调用成功时表示调用结果。 调用失败时无此字段。
type RunModerationAudioResponseBodyResult struct {
	Detail *RunModerationAudioResponseBodyResultDetail `json:"detail,omitempty"`
	// 检测结果是否通过。 block：包含敏感信息，不通过。 pass：不包含敏感信息，通过。 review：需要人工复查

	Suggestion *string `json:"suggestion,omitempty"`
}

func (o RunModerationAudioResponseBodyResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunModerationAudioResponseBodyResult struct{}"
	}

	return strings.Join([]string{"RunModerationAudioResponseBodyResult", string(data)}, " ")
}
