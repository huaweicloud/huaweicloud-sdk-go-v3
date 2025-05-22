package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PublicEndpoints **参数解释**： 公网连接信息。 **取值范围**： 不涉及。
type PublicEndpoints struct {

	// **参数解释**： 公网连接信息。 **取值范围**： 不涉及。
	PublicConnectInfo *string `json:"public_connect_info,omitempty"`

	// **参数解释**： 公网JDBC连接串。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： jdbc:postgresql://<public_connect_info>/<YOUR_DATABASE_name>
	JdbcUrl *string `json:"jdbc_url,omitempty"`
}

func (o PublicEndpoints) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublicEndpoints struct{}"
	}

	return strings.Join([]string{"PublicEndpoints", string(data)}, " ")
}
