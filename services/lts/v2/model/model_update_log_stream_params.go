package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateLogStreamParams struct {

	// 日志存储时间（天）。
	TtlInDays int32 `json:"ttl_in_days"`

	// 标签字段信息
	Tags *[]TagsBody `json:"tags,omitempty"`
}

func (o UpdateLogStreamParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLogStreamParams struct{}"
	}

	return strings.Join([]string{"UpdateLogStreamParams", string(data)}, " ")
}
