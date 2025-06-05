package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CategoryInfo 目录树信息。
type CategoryInfo struct {

	// 主题目录guid。
	Guid *string `json:"guid,omitempty"`

	// 主题目录名称。
	Name *string `json:"name,omitempty"`

	// 主题目录描述信息。
	Description *string `json:"description,omitempty"`

	// 主题目录编码。
	Code *string `json:"code,omitempty"`

	// 主题目录路径。
	Path *string `json:"path,omitempty"`

	// 主题目录别名。
	Alias *string `json:"alias,omitempty"`

	// 主题目录级别。
	Level *string `json:"level,omitempty"`

	// 主题目录排序。
	Ordinal *int32 `json:"ordinal,omitempty"`

	// 主题目录英文名称。
	NameEng *string `json:"name_eng,omitempty"`

	// 主题目录唯一标识名称。
	QualifiedName *string `json:"qualified_name,omitempty"`

	// 资产创建时间。
	CreateTime *int64 `json:"create_time,omitempty"`

	// 资产修改时间。
	UpdateTime *int64 `json:"update_time,omitempty"`

	// 业务意义。
	BusinessMeaning *string `json:"business_meaning,omitempty"`

	// 父级目录guid。
	ParentGuid *string `json:"parent_guid,omitempty"`

	// 主题目录部门。
	Department *string `json:"department,omitempty"`

	// 数据owner所属部门。
	DataOwner *string `json:"data_owner,omitempty"`

	// 数据管家。
	DataSteward *string `json:"data_steward,omitempty"`

	// 数据库。
	Database *[]string `json:"database,omitempty"`

	// obs桶。
	ObsBucket *string `json:"obs_bucket,omitempty"`

	// obs加密桶。
	ObsEncryptBucket *string `json:"obs_encrypt_bucket,omitempty"`

	// 所属空间。
	Workspace *[]string `json:"workspace,omitempty"`

	// 子级目录列表。
	Children *[]CategoryInfo `json:"children,omitempty"`
}

func (o CategoryInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CategoryInfo struct{}"
	}

	return strings.Join([]string{"CategoryInfo", string(data)}, " ")
}
