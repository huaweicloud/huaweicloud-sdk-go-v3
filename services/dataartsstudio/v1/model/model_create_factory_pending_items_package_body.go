package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateFactoryPendingItemsPackageBody struct {

	// 发包人id，可从 查询待发布包列表接口 获取。
	ApplyUserId string `json:"apply_user_id"`

	// 发布包名称
	PackageName string `json:"package_name"`

	// 待发布包id，可从 查询待发布包列表接口 获取。
	PendingItemList []string `json:"pending_item_list"`

	// 审批人id，可通过管理中心-> 查询空间用户信息 接口 获取，且具有管理员或者部署者权限。
	ApproverIds []string `json:"approver_ids"`
}

func (o CreateFactoryPendingItemsPackageBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFactoryPendingItemsPackageBody struct{}"
	}

	return strings.Join([]string{"CreateFactoryPendingItemsPackageBody", string(data)}, " ")
}
