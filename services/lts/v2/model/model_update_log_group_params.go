package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateLogGroupParams 修改日志组的参数。
type UpdateLogGroupParams struct {

	// 日志存储时间 天。 取值范围为 [1, 30]
	TtlInDays int32 `json:"ttl_in_days"`

	// 标签字段信息
	Tags *[]TagsBody `json:"tags,omitempty"`
}

func (o UpdateLogGroupParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLogGroupParams struct{}"
	}

	return strings.Join([]string{"UpdateLogGroupParams", string(data)}, " ")
}
