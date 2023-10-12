package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateOrUpdateStoragePolicyStatementReq 新增或更新存储目录访问权限自定义策略(已存在自定义策略时会对已有策略更新)
type CreateOrUpdateStoragePolicyStatementReq struct {

	// 客户端访问存储可操作的权限合集 允许为空,为空时配置了该策略的用户,通过云办公客户端接入后仅可查看文件列表,不可上传下载 * `PutObject` -  上传、修改、重命名、移动 * `DeleteObject` - 删除 * `GetObject` - 下载 注：PutObject和DeleteObject必须同时设置,不支持仅设置其中一个
	Actions *[]string `json:"actions,omitempty"`

	// 云端访问存储可操作的权限合集,不允许为空 * `PutObject` -  上传、修改、重命名、移动 * `DeleteObject` - 删除 * `GetObject` - 下载           注：PutObject和DeleteObject必须同时设置,不支持仅设置其中一个
	RoamActions []string `json:"roam_actions"`
}

func (o CreateOrUpdateStoragePolicyStatementReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateOrUpdateStoragePolicyStatementReq struct{}"
	}

	return strings.Join([]string{"CreateOrUpdateStoragePolicyStatementReq", string(data)}, " ")
}
