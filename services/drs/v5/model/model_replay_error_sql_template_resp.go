package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReplayErrorSqlTemplateResp 回放异常SQL数据项
type ReplayErrorSqlTemplateResp struct {

	// SQL模板
	SqlTemplate *string `json:"sql_template,omitempty"`

	// SQL模板MD5
	SqlTemplateMd5 *string `json:"sql_template_md5,omitempty"`

	// 目标库昵称
	TargetName *string `json:"target_name,omitempty"`

	// schema名称
	SchemaName *string `json:"schema_name,omitempty"`

	// SQL类型 取值：SELECT, INSERT, DELETE, UPDATE, DDL
	QueryType *string `json:"query_type,omitempty"`

	// 目标库类型 取值：mysql,taurusha,sqlserver,postgresql,ddm,mongodb,awsdocumentdb,hwmongodb,hwpostgresql,oracle,taurus,gaussdb,kafka,mrsKafka,cassandra,dynamo,dws,gaussdbv5,gaussdbv5ha,gaussmongodb,multigaussdbv5,dmq,gaussdbt,gaussdbtha,gaussdb300,elasticsearch,db2,tidb,redis,gaussredis,rediscluster,mariadb,oceanbase,informix,gaussdbv1,gausscassandra,geminidynamo,dds,tauruslite
	TargetType *string `json:"target_type,omitempty"`

	// 归类的SQL数量
	Count *int64 `json:"count,omitempty"`
}

func (o ReplayErrorSqlTemplateResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReplayErrorSqlTemplateResp struct{}"
	}

	return strings.Join([]string{"ReplayErrorSqlTemplateResp", string(data)}, " ")
}
