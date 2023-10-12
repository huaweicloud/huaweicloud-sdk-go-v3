package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeploymentTargetsPrimitiveTypeHolderDeploymentTargets 部署目标信息。
type DeploymentTargetsPrimitiveTypeHolderDeploymentTargets struct {

	// 用户指定资源栈集操作所涉及的区域。  *在DeployStackSet API中，根据用户输入regions和domain_ids列表，以笛卡尔积的形式选择资源栈集中存在的资源栈实例进行部署。如果指定了没有被资源栈集所管理的region，则会报错。*
	Regions []string `json:"regions"`

	// 权限模型是SELF_MANAGED时，用户指定包含本次操作涉及到的租户ID内容。  *在DeployStackSet API中，如果指定该参数，根据用户输入的domain_ids列表和regions列表，以笛卡尔积的形式选择资源栈集中存在的资源栈实例进行部署。如果指定了没有被资源栈集所管理的domain_id，则会报错。*  domain_ids和domain_ids_uri 有且仅有一个存在。
	DomainIds *[]string `json:"domain_ids,omitempty"`

	// 权限模型是SELF_MANAGED时，用户指定包含本次操作涉及到的租户ID内容文件的OBS地址。  内容格式要求每个租户ID以逗号（,）分割，支持换行。当前仅支持csv文件，且文件的编码格式须为UTF-8。文件内容大小不超过100KB。  上传的csv文件应尽量避免Excel操作，以防出现读取内容不一致的问题。推荐使用记事本打开确认内容是否符合预期。  *在DeployStackSet API中，如果指定该参数，根据用户输入的domain_ids_uri文件内容和regions列表，以笛卡尔积的形式选择资源栈集中存在的资源栈实例进行部署。如果内容包含了没有被资源栈集所管理的domain_id，则会报错。*  domain_ids和domain_ids_uri有且仅有一个存在。
	DomainIdsUri *string `json:"domain_ids_uri,omitempty"`
}

func (o DeploymentTargetsPrimitiveTypeHolderDeploymentTargets) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeploymentTargetsPrimitiveTypeHolderDeploymentTargets struct{}"
	}

	return strings.Join([]string{"DeploymentTargetsPrimitiveTypeHolderDeploymentTargets", string(data)}, " ")
}
