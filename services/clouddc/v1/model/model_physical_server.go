package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PhysicalServer 物理服务器信息
type PhysicalServer struct {

	// UUID（Universally Unique Identifier）是一个 128 位的数字，通常以 32 个十六进制数字的形式表示，分为 5 组，用连字符分隔。具体格式如下：  xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx 其中：  每个 x 是一个十六进制数字（0-9 或 a-f）。 5 组的长度分别是：8 位、4 位、4 位、4 位 和 12 位。 为了匹配这种格式的 UUID，可以使用以下正则表达式：  regex ^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$
	Id string `json:"id"`

	// Resource Name Type
	Name string `json:"name"`

	// project id
	ProjectId string `json:"project_id"`

	// domain id
	DomainId string `json:"domain_id"`

	ManageState *ManageState `json:"manage_state"`

	PowerState *PowerState `json:"power_state,omitempty"`

	HealthState *Health `json:"health_state,omitempty"`

	// 上架时间
	OnboardTime string `json:"onboard_time"`

	Location *Location `json:"location"`

	HardwareAttributes *HardwareSummary `json:"hardware_attributes"`

	// 标签
	Tags *[]Tag `json:"tags,omitempty"`

	Error *ErrorStatus `json:"error,omitempty"`
}

func (o PhysicalServer) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PhysicalServer struct{}"
	}

	return strings.Join([]string{"PhysicalServer", string(data)}, " ")
}
