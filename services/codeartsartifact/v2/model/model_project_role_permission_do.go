package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ProjectRolePermissionDo struct {

	// **参数解释**: ID。 **取值范围**: 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**: 角色ID。 **取值范围**: 不涉及。
	RoleId *int32 `json:"role_id,omitempty"`

	// **参数解释**: DEVUC角色ID。 **取值范围**: 不涉及。
	DevucRoleId *string `json:"devuc_role_id,omitempty"`

	// **参数解释**: 项目ID。 **取值范围**: 不涉及。
	ProjectId *string `json:"project_id,omitempty"`

	// **参数解释**: 是否具有权限配置能力。 **取值范围**: - true：能够配置权限。 - false：不能配置权限。
	IsPermissionConfig *bool `json:"is_permission_config,omitempty"`

	// **参数解释**: 能否更改软件包状态。 **取值范围**: - true：能够更改软件包状态。 - false：不能更改软件包状态。
	IsChangePkgStatus *bool `json:"is_change_pkg_status,omitempty"`

	// **参数解释**: 能否进行上传。 **取值范围**: - true：能够进行上传。 - false：不能上传。
	IsUpload *bool `json:"is_upload,omitempty"`

	// **参数解释**: 能否删除和还原测试状态的软件包。 **取值范围**: - true：能够删除和还原测试状态的软件包。 - false：不能删除和还原测试状态的软件包。
	IsDeleteRestoreTestPkg *bool `json:"is_delete_restore_test_pkg,omitempty"`

	// **参数解释**: 能否删除和还原生产状态的软件包。 **取值范围**: - true：能够删除和还原生产状态的软件包。 - false：不能删除和还原生产状态的软件包。
	IsDeleteRestoreProdPkg *bool `json:"is_delete_restore_prod_pkg,omitempty"`

	// **参数解释**: 能否编辑测试状态的软件包。 **取值范围**: - true：能够编辑测试状态的软件包。 - false：不能编辑测试状态的软件包。
	IsEditTestPkg *bool `json:"is_edit_test_pkg,omitempty"`

	// **参数解释**: 能否创建文件夹。 **取值范围**: - true：能够创建目录。 - false：不能创建目录。
	IsMkdir *bool `json:"is_mkdir,omitempty"`

	// **参数解释**: 能否进行下载。 **取值范围**: - true：能够下载。 - false：不能下载。
	IsDownload *bool `json:"is_download,omitempty"`

	// **参数解释**: 能否还原回收站。 **取值范围**: - true：能够在回收站还原所有。 - false：不能在回收站还原所有。
	IsRestoreAll *bool `json:"is_restore_all,omitempty"`

	// **参数解释**: 能否清空回收站。 **取值范围**: - true：能够清空回收站。 - false：不能清空回收站。
	IsEmpty *bool `json:"is_empty,omitempty"`

	// **参数解释**: 创建时间。 **取值范围**: 不涉及。
	CreateTime *int64 `json:"create_time,omitempty"`

	// **参数解释**: 更新时间。 **取值范围**: 不涉及。
	UpdateTime *int64 `json:"update_time,omitempty"`

	// **参数解释**: 权限迁移状态。 **取值范围**: 不涉及。
	Migrated630 *int32 `json:"migrated_630,omitempty"`

	// **参数解释**: 区域。 **取值范围**: 不涉及。
	Region *string `json:"region,omitempty"`

	// **参数解释**: 用户id。 **取值范围**: 不涉及。
	UserId *string `json:"user_id,omitempty"`

	// **参数解释**: 角色。 **取值范围**: 不涉及。
	Roles *string `json:"roles,omitempty"`
}

func (o ProjectRolePermissionDo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProjectRolePermissionDo struct{}"
	}

	return strings.Join([]string{"ProjectRolePermissionDo", string(data)}, " ")
}
