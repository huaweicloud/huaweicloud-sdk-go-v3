package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProxyUpgradeVersionDetail Proxy升级内核版本对象。
type ProxyUpgradeVersionDetail struct {

	// 实例id.
	InstanceId *string `json:"instance_id,omitempty"`

	// 工作流Id。
	WorkflowId *string `json:"workflow_id,omitempty"`

	// agent返回的升级下发状态码，默认返回。
	State *string `json:"state,omitempty"`

	// 错误消息。
	ErrorMessage *string `json:"error_message,omitempty"`
}

func (o ProxyUpgradeVersionDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProxyUpgradeVersionDetail struct{}"
	}

	return strings.Join([]string{"ProxyUpgradeVersionDetail", string(data)}, " ")
}
