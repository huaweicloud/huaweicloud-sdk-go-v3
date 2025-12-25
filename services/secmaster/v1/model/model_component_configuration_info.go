package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ComponentConfigurationInfo struct {

	// 节点配置状态 UN_SAVED 未保存 SAVE_AND_UN_DEPLOY 保存未部署 DEPLOYING正在部署 MOVE_AND_UN_DEPLOY 移除未应用 FAIL_DEPLOY部署失败 DEPLOYED已部署
	ConfigurationStatus *string `json:"configuration_status,omitempty"`

	// 文件参数信息
	List *[]ComponentConfigurationParamVo `json:"list,omitempty"`

	// 节点ID
	NodeId string `json:"node_id"`

	// 节点名
	NodeName string `json:"node_name"`

	// 枚举 节点状态 NORMAL正常 ANOMALIES异常 FAULTS 故障 LOST_CONTACT 失联
	NodeStatus string `json:"node_status"`

	// 节点规格
	Specification string `json:"specification"`
}

func (o ComponentConfigurationInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ComponentConfigurationInfo struct{}"
	}

	return strings.Join([]string{"ComponentConfigurationInfo", string(data)}, " ")
}
