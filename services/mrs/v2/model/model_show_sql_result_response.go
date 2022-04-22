package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowSqlResultResponse struct {

	// SQL的执行id。执行select、show和desc语句时才会生成id，其他操作id为空
	Id *string `json:"id,omitempty"`

	// 错误信息。
	Message *string `json:"message,omitempty"`

	// 执行的SQL语句。
	Statement *string `json:"statement,omitempty"`

	// SQL的执行状态。  - QUEUED - WAITING_FOR_RESOURCES - PLANNING - STARTING - RUNNING - FINISHING - FINISHED - FAILED
	Status *string `json:"status,omitempty"`

	// SQL查询语句的最终结果归档路径。  说明： 只有select的语句才会在将SQL的执行结果转储到result_location中。
	ResultLocation *string `json:"result_location,omitempty"`

	// SQL的执行结果。  说明： 只有非select的语句才会在content中返回结果，如果SQL中没有结果，content为空。
	Content        *[][]string `json:"content,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ShowSqlResultResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSqlResultResponse struct{}"
	}

	return strings.Join([]string{"ShowSqlResultResponse", string(data)}, " ")
}
