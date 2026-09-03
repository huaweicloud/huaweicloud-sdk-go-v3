package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Available 资源可用容量，不考虑资源已分配量，即资源总容量减去故障资源和热备节点的资源
type Available struct {
	Value *Value `json:"value,omitempty"`

	MaxValue *Value `json:"maxValue,omitempty"`

	// UTC时间，格式yyyy-MM-dd'T'HH:mm:ss'Z'。
	Timestamp *string `json:"timestamp,omitempty"`

	// 统计间隔，1s表示1秒，1m表示1分钟，1h为1小时。
	Window *string `json:"window,omitempty"`
}

func (o Available) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Available struct{}"
	}

	return strings.Join([]string{"Available", string(data)}, " ")
}
