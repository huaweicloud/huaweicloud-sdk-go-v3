package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListStoragePolicyStatementResponse Response Object
type ListStoragePolicyStatementResponse struct {

	// 总数
	Count *int32 `json:"count,omitempty"`

	// 支持的访问策略,内置如下三种策略 * `读写` - 上传、编辑、下载 policy_statement_id: DEFAULT_1 action: PutObject、DeleteObject、GetObject * `只读` - 下载 policy_statement_id: DEFAULT_2 action: GetObject * `只写` - 上传、编辑 policy_statement_id: DEFAULT_3 action: PutObject、DeleteObject
	Items          *[]PolicyStatement `json:"items,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListStoragePolicyStatementResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStoragePolicyStatementResponse struct{}"
	}

	return strings.Join([]string{"ListStoragePolicyStatementResponse", string(data)}, " ")
}
