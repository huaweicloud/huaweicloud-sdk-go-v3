package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSparkSqlStateResponse Response Object
type ShowSparkSqlStateResponse struct {

	// **参数解释**：SparkSql作业的ID，用于唯一标识一次SparkSql作业执行。 **取值范围**：采用UUID格式，长度为36个字符，例如：6c98db52-cac2-4ff1-9a91-b7793e95557d。
	StatementId *string `json:"statement_id,omitempty"`

	// **参数解释**：SparkSql作业的执行状态。 **取值范围**：   - QUEUED：排队中。   - RUNNING：运行中。   - CANCELING：取消中。   - CANCELED：已取消。   - FAILED：运行失败。   - QUEUED_TIMEOUT：排队超时。   - RUNNING_TIMEOUT：运行超时。   - SUCCEED：运行成功。
	State          *string `json:"state,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowSparkSqlStateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSparkSqlStateResponse struct{}"
	}

	return strings.Join([]string{"ShowSparkSqlStateResponse", string(data)}, " ")
}
