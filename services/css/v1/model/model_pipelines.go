package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Pipelines struct {

	// 配置文件名称。
	Name *string `json:"name,omitempty" xml:"name"`

	// pipeline状态。
	Status *string `json:"status,omitempty" xml:"status"`

	// 是否开启常驻。
	KeepAlive *bool `json:"keepAlive,omitempty" xml:"keepAlive"`

	// 事件只有在“工作中”状态才可以实时查看（需要手动刷新），“已停止”状态请到output端查看迁移数据量。
	Events *string `json:"events,omitempty" xml:"events"`

	// 更新时间。
	UpdateAt *string `json:"updateAt,omitempty" xml:"updateAt"`
}

func (o Pipelines) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Pipelines struct{}"
	}

	return strings.Join([]string{"Pipelines", string(data)}, " ")
}
