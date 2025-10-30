package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAgentDaemonsetInfoResponse Response Object
type ListAgentDaemonsetInfoResponse struct {

	// **参数解释**: 数据列表 **取值范围**: 取值0-200
	DataList *[]ClusterInfoResponse `json:"data_list,omitempty"`

	// **参数解释** 已接入集群总数 **取值范围** 取值0-65535
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释** 待升级集群总数 **取值范围** 取值0-65535
	UpgradefulNum *int32 `json:"upgradeful_num,omitempty"`

	// **参数解释** 运行异常集群总数 **取值范围** 取值0-65535
	ErrRunningNum *int32 `json:"err_running_num,omitempty"`

	// **参数解释** 接入异常集群总数 **取值范围** 取值0-65535
	ErrAccessNum   *int32 `json:"err_access_num,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListAgentDaemonsetInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAgentDaemonsetInfoResponse struct{}"
	}

	return strings.Join([]string{"ListAgentDaemonsetInfoResponse", string(data)}, " ")
}
