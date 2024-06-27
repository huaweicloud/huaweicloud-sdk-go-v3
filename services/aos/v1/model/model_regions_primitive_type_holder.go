package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RegionsPrimitiveTypeHolder struct {

	// 用户指定资源栈集操作所涉及的区域。  *在DeployStackSet API中，根据用户输入regions和domain_ids列表，以笛卡尔积的形式选择资源栈集中存在的资源栈实例进行部署。如果指定了没有被资源栈集所管理的region，则会报错。*
	Regions []string `json:"regions"`
}

func (o RegionsPrimitiveTypeHolder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegionsPrimitiveTypeHolder struct{}"
	}

	return strings.Join([]string{"RegionsPrimitiveTypeHolder", string(data)}, " ")
}
