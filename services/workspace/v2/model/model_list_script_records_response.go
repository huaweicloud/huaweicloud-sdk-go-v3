package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListScriptRecordsResponse Response Object
type ListScriptRecordsResponse struct {

	// 总数。
	Count *int32 `json:"count,omitempty"`

	// 脚本执行记录列表。
	ScriptRecords  *[]ScriptRecordSimpleInfo `json:"script_records,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ListScriptRecordsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScriptRecordsResponse struct{}"
	}

	return strings.Join([]string{"ListScriptRecordsResponse", string(data)}, " ")
}
