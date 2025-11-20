package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MasterEipRequestSpecSpec **参数解释**： 待绑定的弹性IP配置属性 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
type MasterEipRequestSpecSpec struct {

	// **参数解释**： 弹性网卡ID，绑定时必选，解绑时该字段无效 > 获取方式:在控制台的“服务列表”中，单击“网络 > 弹性公网IP EIP”，单击弹性公网IP，在详情页的“基本信息”页签下找到“ID”字段复制即可。  **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Id *string `json:"id,omitempty"`
}

func (o MasterEipRequestSpecSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MasterEipRequestSpecSpec struct{}"
	}

	return strings.Join([]string{"MasterEipRequestSpecSpec", string(data)}, " ")
}
