package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAuraSqlStatementResultResponse Response Object
type ShowAuraSqlStatementResultResponse struct {

	// **参数解释**：SQL执行结果，可见下方响应示例。 **取值范围**： ## StatementResponse - **status**（Integer）：请求整体执行状态。   - 0：成功。   - -1：失败。   - 1： 执行中。   - 2：执行完成但无结果集。   - 3：排队等待中。   - 4：语句未正常执行。 - **session_id**（String，UUID）：当前执行会话唯一标识。 - **statement_id**（String，UUID）：本次语句执行唯一标识。 - **results**（Array of StatementResult objects）：每条SQL执行结果。 --- ## StatementResult - **status**（String）：SQL执行状态。   - PGRES_TUPLES_OK：查询成功并含结果集。   - PGRES_COMMAND_OK：查询成功并含结果集。   - PGRES_FATAL_ERROR：执行失败。 - **statement_id**（String）：语句ID。 - **num_rows**（Integer）：查询结果总行数。 - **row_count**（Integer）：当前页实际返回行数。 - **page_no**（Integer）：当前页码。 - **page_count**（Integer）：总页数。 - **err_code**（String）：错误码，0表示无错误，其他数值参见错误信息。 - **sql_state**（String）：SQL状态码。 - **message**（String）：执行信息或错误详情。 - **result_set**（Object，type: StatementResultSet）：仅查询成功时含有效数据，否则为空结构。 --- ## StatementResultSet - **columns**（Array of RowType objects）：列元数据列表。 - **rows**（Array of String arrays）：实际数据行，每行顺序与columns对应。 --- ## RowType - **name**（String）：列名。 - **table_id**（Integer）：表的ID。 - **column_id**（Integer）：列的ID。 - **format**（Integer）：格式。 - **type**（Integer）：PG类型OID。 - **size**（Integer）：大小。 - **type_mod**（Integer）：typemod。
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowAuraSqlStatementResultResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAuraSqlStatementResultResponse struct{}"
	}

	return strings.Join([]string{"ShowAuraSqlStatementResultResponse", string(data)}, " ")
}
