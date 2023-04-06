package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// server字段数据结构说明。
type RespServer struct {

	// 弹性云服务器的网络属性。
	Addresses map[string][]RespAddr `json:"addresses"`

	// 弹性云服务器创建时间。
	Created string `json:"created"`

	Flavor *RespFlavor `json:"flavor"`

	// 弹性云服务器ID，格式为UUID。
	Id string `json:"id"`

	// 弹性云服务器名称。
	Name string `json:"name"`

	// 弹性云服务器状态。  取值范围：ACTIVE、BUILD、DELETED、ERROR、HARD_REBOOT、MIGRATING、PASSWORD、PAUSED、REBOOT、REBUILD、RESIZE、REVERT_RESIZE、SHUTOFF、SHELVED、SHELVED_OFFLOADED、SOFT_DELETED、SUSPENDED、VERIFY_RESIZE
	Status string `json:"status"`

	// 弹性云服务器所属租户ID，格式为UUID。
	TenantId string `json:"tenant_id"`

	// 弹性云服务器更新时间。
	Updated string `json:"updated"`

	// 创建弹性云服务器的用户ID，格式为UUID。
	UserId string `json:"user_id"`

	// 弹性云服务器当前任务的状态。
	TaskState string `json:"task_state"`

	Image *RespImage `json:"image"`

	Metadata *RespMetadata `json:"metadata"`
}

func (o RespServer) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RespServer struct{}"
	}

	return strings.Join([]string{"RespServer", string(data)}, " ")
}
