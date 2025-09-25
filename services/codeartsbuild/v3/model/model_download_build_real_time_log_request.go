package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DownloadBuildRealTimeLogRequest Request Object
type DownloadBuildRealTimeLogRequest struct {

	// 构建的任务ID； 编辑构建任务时，浏览器URL末尾的32位数字、字母组合的字符串。
	JobId string `json:"job_id"`

	// 构建任务的构建编号，从1开始，每次构建递增1
	BuildNo int32 `json:"build_no"`

	// **参数解释**： 起始偏移量，表示从此开始查询。 **约束限制**： 不涉及。 **取值范围**： 只能使用数字，大于等于0。
	StartOffset *int32 `json:"start_offset,omitempty"`

	// **参数解释**： 结束偏移量，表示查询到此结束。 **约束限制**： 不涉及。 **取值范围**： 只能使用数字，大于等于0。
	EndOffset *int32 `json:"end_offset,omitempty"`

	// **参数解释**： 排序方式。 **约束限制**： 不涉及。 **取值范围**： AES|DESC。
	Sort *string `json:"sort,omitempty"`

	// **参数解释**： 可控制返回内容长度，固定1000。 **约束限制**： 不涉及。 **取值范围**： 只能使用数字，大于等于0。 **默认取值**： 1000
	Size string `json:"size"`
}

func (o DownloadBuildRealTimeLogRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadBuildRealTimeLogRequest struct{}"
	}

	return strings.Join([]string{"DownloadBuildRealTimeLogRequest", string(data)}, " ")
}
