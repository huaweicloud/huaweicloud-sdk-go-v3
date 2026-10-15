package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSparkSqlResItem **参数解释**：SparkSql作业的响应参数详情。
type ListSparkSqlResItem struct {

	// **参数解释**：SparkSql作业的ID，用于唯一标识一次SparkSql作业执行。 **取值范围**：采用UUID格式，长度为36个字符，例如：80ceaaff-3cfc-4162-a56f-70031ea4fa91。
	StatementId string `json:"statement_id"`

	// **参数解释**：SparkSql作业事务ID，用于防止重复提交。 **取值范围**：采用UUID格式，长度为36个字符，例如：xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx。
	ClientToken *string `json:"client_token,omitempty"`

	// **参数解释**：端点名称，用于指定SparkSql执行环境。 **取值范围**：只能由英文小写字母、数字及中划线组成，以英文小写字母开头，以英文小写字母或数字结尾，且长度为1~63个字符。
	EndpointName string `json:"endpoint_name"`

	CatalogContext *SparkSqlCatalogContextResponse `json:"catalog_context"`

	// **参数解释**：用户SQL语句，用于执行数据查询、数据操作等任务。 **取值范围**：长度不超过500000个字符。
	Statement string `json:"statement"`

	// **参数解释**：用户SQL内容的占位符参数列表，用于动态替换SQL中的参数。
	Parameters *[]SparkSqlParameter `json:"parameters,omitempty"`

	// **参数解释**：SparkSql作业的状态，用于标识作业的执行状态。 **取值范围**： - QUEUED：排队中。 - RUNNING：运行中。 - CANCELING：取消中。 - CANCELED：已取消。 - FAILED：运行失败。 - QUEUED_TIMEOUT：排队超时。 - RUNNING_TIMEOUT：运行超时。 - SUCCEED：运行成功。
	State string `json:"state"`

	// **参数解释**：SQL作业类型，用于标识作业的类型。 **取值范围**： - DDL：创建修改删除元数据类型的作业、DESC/SHOW等语句。 - DCL：权限授权与回收类型的作业。 - DQL：查询语句SELECT。 - DML：向表追加、删除、更新新数据类型的作业。
	StatementType string `json:"statement_type"`

	Error *SparkSqlErrorDto `json:"error,omitempty"`

	Result *SparkSqlResultResponse `json:"result,omitempty"`

	MetricStatistics *SparkSqlMetricStatisticsResponse `json:"metric_statistics,omitempty"`

	// **参数解释**：日志OBS归档路径，用于存储作业执行日志。 **取值范围**：采用OBS路径格式，例如：obs://bucket/aidatalake/workspace_xxx/spark/endpoint_xxx/jobs/logs/2026_04_27/sql-engine-{statement_id}/spark.log。
	LogUrl *string `json:"log_url,omitempty"`

	// **参数解释**：作业创建时间，采用unix时间戳格式。 **取值范围**：单位为毫秒。
	CreateTime int64 `json:"create_time"`

	// **参数解释**：作业开始运行时间，采用unix时间戳格式。 **取值范围**：单位为毫秒。
	StartTime *int64 `json:"start_time,omitempty"`

	// **参数解释**：作业结束时间，采用unix时间戳格式。 **取值范围**：单位为毫秒。
	EndTime *int64 `json:"end_time,omitempty"`

	CreateUser *SparkCreateUser `json:"create_user"`
}

func (o ListSparkSqlResItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSparkSqlResItem struct{}"
	}

	return strings.Join([]string{"ListSparkSqlResItem", string(data)}, " ")
}
