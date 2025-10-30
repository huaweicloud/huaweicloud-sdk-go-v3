package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UninstallAgentsRequestInfo struct {

	// **参数解释**： 服务器ID列表 **约束限制**： 不涉及 **取值范围**： 列表长度为1-200 **默认取值**： 不涉及
	HostIdList []string `json:"host_id_list"`
}

func (o UninstallAgentsRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UninstallAgentsRequestInfo struct{}"
	}

	return strings.Join([]string{"UninstallAgentsRequestInfo", string(data)}, " ")
}
