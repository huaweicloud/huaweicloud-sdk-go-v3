package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WarehouseApp 应用仓库中的应用详细信息。
type WarehouseApp struct {

	// 应用的记录id。
	Id *string `json:"id,omitempty"`

	// 应用id。
	AppId *string `json:"app_id,omitempty"`

	// 租户id。
	TenantId *string `json:"tenant_id,omitempty"`

	// 应用名称。
	AppName *string `json:"app_name,omitempty"`

	AppCategory *AppCategoryEnum `json:"app_category,omitempty"`

	OsType *OsTypeEnum `json:"os_type,omitempty"`

	// 版本号。
	VersionId *string `json:"version_id,omitempty"`

	// 版本名称。
	VersionName *string `json:"version_name,omitempty"`

	// 应用文件的存放路径。
	AppfileStorePath *string `json:"appfile_store_path,omitempty"`

	// 应用文件的大小，以KB为单位。
	AppFileSize *string `json:"app_file_size,omitempty"`

	// 应用描述。
	AppDescription *string `json:"app_description,omitempty"`

	// 应用文件的存放路径。
	AppiconStorePath *string `json:"appicon_store_path,omitempty"`

	// 应用创建时间。
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 应用修改时间。
	ModifyTime *sdktime.SdkTime `json:"modify_time,omitempty"`

	// 应用审核时间。
	VerifyTime *sdktime.SdkTime `json:"verify_time,omitempty"`

	VerifyStatus *VerifyStatusEnum `json:"verify_status,omitempty"`

	// 审核的评论意见。
	VerifyComment *string `json:"verify_comment,omitempty"`

	// app的图标文件。
	AppIcon *string `json:"app_icon,omitempty"`
}

func (o WarehouseApp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WarehouseApp struct{}"
	}

	return strings.Join([]string{"WarehouseApp", string(data)}, " ")
}
