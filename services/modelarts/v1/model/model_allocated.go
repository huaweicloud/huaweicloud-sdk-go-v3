package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Allocated 资源已分配量。
type Allocated struct {
	Value *Value `json:"value,omitempty"`

	// UTC时间，格式yyyy-MM-dd'T'HH:mm:ss'Z'。
	Timestamp *string `json:"timestamp,omitempty"`

	// 统计间隔，1s表示1秒，1m表示1分钟，1h为1小时。
	Window *string `json:"window,omitempty"`
}

func (o Allocated) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Allocated struct{}"
	}

	return strings.Join([]string{"Allocated", string(data)}, " ")
}
