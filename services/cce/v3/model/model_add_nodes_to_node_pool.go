package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddNodesToNodePool **参数解释**： 自定义节点池纳管节点参数。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
type AddNodesToNodePool struct {

	// **参数解释**： 服务器ID，从ECS/BMS控制台获取。 > 获取方式：在控制台的“服务列表”中，单击“计算 > 弹性云服务器 ECS/裸金属服务器 BMS”，单击服务器的名称，在服务器详情页的“基本信息”页签下找到“ID”字段复制即可。  **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	ServerID string `json:"serverID"`
}

func (o AddNodesToNodePool) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddNodesToNodePool struct{}"
	}

	return strings.Join([]string{"AddNodesToNodePool", string(data)}, " ")
}
