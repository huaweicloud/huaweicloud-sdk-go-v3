package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateOrUpdateStoragePolicyStatementResponse Response Object
type CreateOrUpdateStoragePolicyStatementResponse struct {

	// 支持的访问策略，内置如下四种策略: * `DEFAULT_1`：`客户端访问存储` - 上传、下载; `云端访问存储` - 读写   - action: PutObject、DeleteObject、GetObject   - roam_action: PutObject、DeleteObject、GetObject * `DEFAULT_2`：`客户端访问存储` - 下载; `云端访问存储` - 读写   - action: GetObject   - roam_action: PutObject、DeleteObject、GetObject * `DEFAULT_3`：`客户端访问存储` - 上传; `云端访问存储` - 读写   - action: PutObject、DeleteObject   - roam_action: PutObject、DeleteObject、GetObject * `DEFAULT_4`：`客户端访问存储` - 仅可查看列表,不允许上传下载; `云端访问存储` - 只读   - action:    - roam_action: GetObject
	PolicyStatementId *string `json:"policy_statement_id,omitempty"`

	// 客户端访问存储可操作的权限合集 * `PutObject` -  上传、修改、重命名、移动 * `GetObject` - 下载 * `DeleteObject` - 删除
	Actions *[]string `json:"actions,omitempty"`

	// 云端访问存储可操作的权限合集 * `PutObject` -  上传、修改、重命名、移动 * `GetObject` - 下载 * `DeleteObject` - 删除
	RoamActions    *[]string `json:"roam_actions,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o CreateOrUpdateStoragePolicyStatementResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateOrUpdateStoragePolicyStatementResponse struct{}"
	}

	return strings.Join([]string{"CreateOrUpdateStoragePolicyStatementResponse", string(data)}, " ")
}
