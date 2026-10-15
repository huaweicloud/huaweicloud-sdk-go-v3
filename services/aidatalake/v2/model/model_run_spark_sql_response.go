package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunSparkSqlResponse Response Object
type RunSparkSqlResponse struct {

	// **参数解释**：作业ID，用于唯一标识本次执行的SparkSql作业。可通过此ID查询作业状态和执行结果。 **取值范围**：采用UUID格式，长度为36个字符，例如：6c98db52-cac2-4ff1-9a91-b7793e95557d。
	StatementId    *string `json:"statement_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RunSparkSqlResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunSparkSqlResponse struct{}"
	}

	return strings.Join([]string{"RunSparkSqlResponse", string(data)}, " ")
}
