package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Endpoints **参数解释**： 集群的内网连接信息。 **取值范围**： 不涉及。
type Endpoints struct {

	// **参数解释**： 内网连接信息。 **取值范围**： 不涉及。
	ConnectInfo *string `json:"connect_info,omitempty"`

	// **参数解释**： 内网JDBC URL。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： jdbc:postgresql://<connect_info>/<YOUR_DATABASE_NAME>
	JdbcUrl *string `json:"jdbc_url,omitempty"`
}

func (o Endpoints) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Endpoints struct{}"
	}

	return strings.Join([]string{"Endpoints", string(data)}, " ")
}
