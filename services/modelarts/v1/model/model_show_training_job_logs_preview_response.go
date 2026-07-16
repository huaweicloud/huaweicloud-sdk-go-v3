package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTrainingJobLogsPreviewResponse Response Object
type ShowTrainingJobLogsPreviewResponse struct {

	// 日志内容。如果日志大小没有超过上限（n兆）则返回全部内容，如果日志超过了上限（n兆）则返回最新的n兆的日志。2022/03/01 00:00:00 (GMT+08:00)后，此参数名称由“context”改为“content”。
	Content *string `json:"content,omitempty"`

	// 当前返回的日志大小（单位：字节）。最大为5兆。
	CurrentSize *int32 `json:"current_size,omitempty"`

	// 完整的日志大小（单位：字节）。
	FullSize       *int32 `json:"full_size,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowTrainingJobLogsPreviewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTrainingJobLogsPreviewResponse struct{}"
	}

	return strings.Join([]string{"ShowTrainingJobLogsPreviewResponse", string(data)}, " ")
}
