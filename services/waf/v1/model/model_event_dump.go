package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EventDump struct {

	// 文件来源
	Source *string `json:"source,omitempty"`

	// 文件状态
	State *string `json:"state,omitempty"`

	// 事件描述
	Description *string `json:"description,omitempty"`

	// 文件id
	Id *string `json:"id,omitempty"`

	// 文件名
	Filename *string `json:"filename,omitempty"`

	// 文件obs名
	Obsname *string `json:"obsname,omitempty"`

	// 租户id
	Tenantid *string `json:"tenantid,omitempty"`

	// 统计开始时间
	Start *int64 `json:"start,omitempty"`

	// 统计截止时间
	End *int64 `json:"end,omitempty"`

	// 总计事件数
	Total *int64 `json:"total,omitempty"`

	// 链接
	Url *string `json:"url,omitempty"`

	// 更新url时间戳
	Urltimestamp *int64 `json:"urltimestamp,omitempty"`

	// 文件生成时间戳
	Timestamp *int64 `json:"timestamp,omitempty"`
}

func (o EventDump) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventDump struct{}"
	}

	return strings.Join([]string{"EventDump", string(data)}, " ")
}
