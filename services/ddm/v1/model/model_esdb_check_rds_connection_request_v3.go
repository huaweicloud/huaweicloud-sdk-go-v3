package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EsdbCheckRdsConnectionRequestV3 struct {

	// **参数解释**：  内网地址以及端口号。  格式为xx.xx.xx.xx:xx。  **参数范围**：  不涉及。
	Endpoint string `json:"endpoint"`

	// **参数解释**：  实例账号名称。  **参数范围**：  不涉及。
	Username string `json:"username"`

	// **参数解释**：  实例账号密码。  **约束限制**：  - 至少包含三种字符组合：小写字母、大写字母、数字、特殊字符 ~ ! @ # % ^ * - _ + ? - 不能使用简单、强度不够、容易被猜测的密码。 - 不能与账号名称或者倒序的账号名称相同。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	Password string `json:"password"`

	// **参数解释**：  rds实例ID，此参数是实例的唯一标识。  **约束限制**：  不涉及。  **取值范围**：  只能由英文字母、数字组成，后缀为in01，长度为36个字符。  **默认取值**：  不涉及。
	RdsInstanceId string `json:"rds_instance_id"`
}

func (o EsdbCheckRdsConnectionRequestV3) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EsdbCheckRdsConnectionRequestV3 struct{}"
	}

	return strings.Join([]string{"EsdbCheckRdsConnectionRequestV3", string(data)}, " ")
}
