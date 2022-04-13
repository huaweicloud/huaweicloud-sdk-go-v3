package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RunImageBatchModerationResponse struct {
	// 调用成功时表示调用结果。 调用失败时无此字段。

	Result         *[]ImageBatchModerationResultBody `json:"result,omitempty"`
	HttpStatusCode int                               `json:"-"`
}

func (o RunImageBatchModerationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunImageBatchModerationResponse struct{}"
	}

	return strings.Join([]string{"RunImageBatchModerationResponse", string(data)}, " ")
}
