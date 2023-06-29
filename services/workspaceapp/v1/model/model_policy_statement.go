package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PolicyStatement 支持的访问策略
type PolicyStatement struct {

	// 支持的访问策略,内置如下三种策略 * `读写` - 上传、编辑、下载 policy_statement_id: DEFAULT_1 * `只读` - 下载 policy_statement_id: DEFAULT_2 * `只写` - 上传、编辑 policy_statement_id: DEFAULT_3
	PolicyStatementId *string `json:"policy_statement_id,omitempty"`

	// 可以进行操作的权限合集 * `PutObject` -  上传、修改、重命名、移动 * `GetObject` - 下载 * `DeleteObject` - 删除
	Actions *[]string `json:"actions,omitempty"`
}

func (o PolicyStatement) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyStatement struct{}"
	}

	return strings.Join([]string{"PolicyStatement", string(data)}, " ")
}
