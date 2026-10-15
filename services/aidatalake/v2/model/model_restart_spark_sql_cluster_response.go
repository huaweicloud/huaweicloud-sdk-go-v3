package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestartSparkSqlClusterResponse Response Object
type RestartSparkSqlClusterResponse struct {

	// **参数解释**：重启操作ID，用于标识本次重启操作。 **取值范围**：只能由英文字母、数字及中划线组成，且长度为1~64个字符。
	OperationId    *string `json:"operation_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RestartSparkSqlClusterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestartSparkSqlClusterResponse struct{}"
	}

	return strings.Join([]string{"RestartSparkSqlClusterResponse", string(data)}, " ")
}
