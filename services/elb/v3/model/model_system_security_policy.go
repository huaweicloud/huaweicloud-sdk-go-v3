package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SystemSecurityPolicy struct {

	// **参数解释**：系统安全策略的名称。  **取值范围**：不涉及
	Name string `json:"name"`

	// **参数解释**：系统安全策略的TLS协议列表。  **取值范围**：不涉及
	Protocols string `json:"protocols"`

	// **参数解释**：系统安全策略的加密套件列表。  **取值范围**：不涉及
	Ciphers string `json:"ciphers"`

	// **参数解释**：项目ID。获取方式请参见[获取项目ID](elb_fl_0008.xml)。  **取值范围**：长度为32个字符，由小写字母和数字组成。
	ProjectId string `json:"project_id"`
}

func (o SystemSecurityPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SystemSecurityPolicy struct{}"
	}

	return strings.Join([]string{"SystemSecurityPolicy", string(data)}, " ")
}
