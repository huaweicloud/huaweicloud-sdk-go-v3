package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestartSparkSqlClusterRequestBody **参数解释**：重启SparkSql集群的请求参数。
type RestartSparkSqlClusterRequestBody struct {
	RestartStrategy *SparkSqlRestartStrategy `json:"restart_strategy"`
}

func (o RestartSparkSqlClusterRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestartSparkSqlClusterRequestBody struct{}"
	}

	return strings.Join([]string{"RestartSparkSqlClusterRequestBody", string(data)}, " ")
}
