package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListReleasePackagesRespData struct {

	// 申请时间，13位时间戳
	ApplyTimestamp *int64 `json:"apply_timestamp,omitempty"`

	// 申请人id
	ApplyUserId *string `json:"apply_user_id,omitempty"`

	// 申请人名称
	ApplyUserName *string `json:"apply_user_name,omitempty"`

	// 是否删除，0:没有删除，1:删除
	DeleteFlag *int32 `json:"delete_flag,omitempty"`

	// 发布状态，1:待审批,2:成功,3:失败, 5:发布中
	DeployStatus *int32 `json:"deploy_status,omitempty"`

	// 发布时间，13位时间戳
	DeployTimestamp *int64 `json:"deploy_timestamp,omitempty"`

	// 发布人id
	DeployUserId *string `json:"deploy_user_id,omitempty"`

	// 发布人名称
	DeployUserName *string `json:"deploy_user_name,omitempty"`

	// 发布包审批信息
	PackageApprovers []ListReleasePackagesRespPackageApprovers `json:"package_approvers"`

	// 发布包id
	PackageId *string `json:"package_id,omitempty"`

	// 发布包名称
	PackageName *string `json:"package_name,omitempty"`

	// 项目ID和空间信息，以 项目ID-workspace-空间ID 拼接。
	ProjectId *string `json:"project_id,omitempty"`

	// 发布包所在空间ID
	WorkspaceId *string `json:"workspace_id,omitempty"`
}

func (o ListReleasePackagesRespData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListReleasePackagesRespData struct{}"
	}

	return strings.Join([]string{"ListReleasePackagesRespData", string(data)}, " ")
}
