package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateTableRequestBody struct {

	// 表别名
	TableAlias *string `json:"table_alias,omitempty"`

	// 表别名
	TableAliasEn *string `json:"table_alias_en,omitempty"`

	// 表别名
	TableAliasFr *string `json:"table_alias_fr,omitempty"`

	// 表相关描述
	Description *string `json:"description,omitempty"`

	// 表相关描述
	DescriptionEn *string `json:"description_en,omitempty"`

	// 表相关描述
	DescriptionFr *string `json:"description_fr,omitempty"`

	// 目录分组
	Directory *string `json:"directory,omitempty"`

	// 目录分组
	DirectoryEn *string `json:"directory_en,omitempty"`

	// 目录分组
	DirectoryFr *string `json:"directory_fr,omitempty"`

	StorageSetting *IsapTableStorageSettingDto `json:"storage_setting,omitempty"`

	DisplaySetting *IsapTableDisplaySettingDto `json:"display_setting,omitempty"`
}

func (o UpdateTableRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTableRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateTableRequestBody", string(data)}, " ")
}
