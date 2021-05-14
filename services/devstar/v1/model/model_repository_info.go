package model

import (
	"encoding/json"

	"strings"
)

type RepositoryInfo struct {
	// 代码仓的名称。

	Name *string `json:"name,omitempty"`
	// 代码仓描述。

	Description *string `json:"description,omitempty"`
	// 项目id。

	ProjectId *string `json:"project_id,omitempty"`
	// 区域id。

	RegionId *string `json:"region_id,omitempty"`
	// 根据编程语言生成gitignore文件。

	Gitignore *string `json:"gitignore,omitempty"`
	// 是否允许项目内成员访问仓库： - 0：不允许 - 1：允许

	MemberPermission *int32 `json:"member_permission,omitempty"`
	// 是否允许生成README文件： - 0：不允许 - 1：允许

	ReadmePermission *int32 `json:"readme_permission,omitempty"`
	// 是否公开： - 0：私有 - 20：公开只读

	VisibilityLevel *int32 `json:"visibility_level,omitempty"`
	//  开源许可证id （0:默认）。

	LicenseId *int32 `json:"license_id,omitempty"`
}

func (o RepositoryInfo) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RepositoryInfo struct{}"
	}

	return strings.Join([]string{"RepositoryInfo", string(data)}, " ")
}
