package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

type ExtensionVersionSnake struct {

	// 插件版本id
	Id *string `json:"id,omitempty"`

	// 插件版本号
	Version *string `json:"version,omitempty"`

	// 版本排序
	VersionRanking *int64 `json:"version_ranking,omitempty"`

	// 插件版本状态 - INIT 待发布 - VALIDATING 审核中 - REJECTED 审核拒绝 - PUBLISHED 插件上架 - OFFLINE 插件下线 - ABANDONED 废弃 - GRAY_INIT 灰度审核 - GRAYED 灰度发布 - GRAY_REJECTED 灰度拒绝
	Status *ExtensionVersionSnakeStatus `json:"status,omitempty"`

	// 插件状态 - INIT 待发布 - VALIDATING 审核中 - REJECTED 审核拒绝 - PUBLISHED 插件上架 - OFFLINE 插件下线 - ABANDONED 废弃 - GRAY_INIT 灰度审核 - GRAYED 灰度发布 - GRAY_REJECTED 灰度拒绝
	VersionStatus *ExtensionVersionSnakeVersionStatus `json:"version_status,omitempty"`

	// 资源文件url
	AssetUri *string `json:"asset_uri,omitempty"`

	// 更新时间
	LastUpdated *sdktime.SdkTime `json:"last_updated,omitempty"`

	// 插件文件集合
	Files *[]ExtensionFileSnake `json:"files,omitempty"`

	// 插件审核信息
	ValidateMessage *string `json:"validate_message,omitempty"`

	// 插件审核状态 - NONE 无 - UPLOADING 上传中 - VALIDATING 系统审核 - OFFLINING 用户申请下线 - ONLINING 用户申请上线 - UMS_VALIDATING 发布商审核中
	VersionValidateStatus *ExtensionVersionSnakeVersionValidateStatus `json:"version_validate_status,omitempty"`

	// 插件展示名称
	DisplayName *string `json:"display_name,omitempty"`

	// 插件描述
	Description *string `json:"description,omitempty"`

	// 插件支持ide版本
	MinIdeVersion *string `json:"min_ide_version,omitempty"`

	// 支持的最大版本
	MaxIdeVersion *string `json:"max_ide_version,omitempty"`

	// 发布时间
	VersionDate *sdktime.SdkTime `json:"version_date,omitempty"`

	// 是否预览
	Preview *bool `json:"preview,omitempty"`

	// 包含插件列表
	ExtensionPack *string `json:"extension_pack,omitempty"`

	// 依赖插件列表
	ExtensionDependencies *string `json:"extension_dependencies,omitempty"`

	// 创建时间
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	// 支持的ide编码
	SupportIde *int32 `json:"support_ide,omitempty"`

	// 插件包源码仓
	RepoUrl *string `json:"repo_url,omitempty"`

	// 帮助页面
	HelpPage *string `json:"help_page,omitempty"`

	// 产品首页
	Website *string `json:"website,omitempty"`

	// 问题链接
	IssueLink *string `json:"issue_link,omitempty"`

	// 插件大小
	AssetSize *int64 `json:"asset_size,omitempty"`

	// 依赖插件
	Depends *[]string `json:"depends,omitempty"`

	// CodeArtsIDEOnline插件版本参数
	PropertyList *[]CodeArtsIdeOnlineExtensionVersionProperty `json:"property_list,omitempty"`

	// 版本发布者
	Uploader *string `json:"uploader,omitempty"`

	// 插件id
	ExtensionId *string `json:"extension_id,omitempty"`
}

func (o ExtensionVersionSnake) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExtensionVersionSnake struct{}"
	}

	return strings.Join([]string{"ExtensionVersionSnake", string(data)}, " ")
}

type ExtensionVersionSnakeStatus struct {
	value string
}

type ExtensionVersionSnakeStatusEnum struct {
	INIT          ExtensionVersionSnakeStatus
	VALIDATING    ExtensionVersionSnakeStatus
	REJECTED      ExtensionVersionSnakeStatus
	PUBLISHED     ExtensionVersionSnakeStatus
	OFFLINE       ExtensionVersionSnakeStatus
	ABANDONED     ExtensionVersionSnakeStatus
	GRAY_INIT     ExtensionVersionSnakeStatus
	GRAYED        ExtensionVersionSnakeStatus
	GRAY_REJECTED ExtensionVersionSnakeStatus
}

func GetExtensionVersionSnakeStatusEnum() ExtensionVersionSnakeStatusEnum {
	return ExtensionVersionSnakeStatusEnum{
		INIT: ExtensionVersionSnakeStatus{
			value: "INIT",
		},
		VALIDATING: ExtensionVersionSnakeStatus{
			value: "VALIDATING",
		},
		REJECTED: ExtensionVersionSnakeStatus{
			value: "REJECTED",
		},
		PUBLISHED: ExtensionVersionSnakeStatus{
			value: "PUBLISHED",
		},
		OFFLINE: ExtensionVersionSnakeStatus{
			value: "OFFLINE",
		},
		ABANDONED: ExtensionVersionSnakeStatus{
			value: "ABANDONED",
		},
		GRAY_INIT: ExtensionVersionSnakeStatus{
			value: "GRAY_INIT",
		},
		GRAYED: ExtensionVersionSnakeStatus{
			value: "GRAYED",
		},
		GRAY_REJECTED: ExtensionVersionSnakeStatus{
			value: "GRAY_REJECTED",
		},
	}
}

func (c ExtensionVersionSnakeStatus) Value() string {
	return c.value
}

func (c ExtensionVersionSnakeStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExtensionVersionSnakeStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type ExtensionVersionSnakeVersionStatus struct {
	value string
}

type ExtensionVersionSnakeVersionStatusEnum struct {
	INIT          ExtensionVersionSnakeVersionStatus
	VALIDATING    ExtensionVersionSnakeVersionStatus
	REJECTED      ExtensionVersionSnakeVersionStatus
	PUBLISHED     ExtensionVersionSnakeVersionStatus
	OFFLINE       ExtensionVersionSnakeVersionStatus
	ABANDONED     ExtensionVersionSnakeVersionStatus
	GRAY_INIT     ExtensionVersionSnakeVersionStatus
	GRAYED        ExtensionVersionSnakeVersionStatus
	GRAY_REJECTED ExtensionVersionSnakeVersionStatus
}

func GetExtensionVersionSnakeVersionStatusEnum() ExtensionVersionSnakeVersionStatusEnum {
	return ExtensionVersionSnakeVersionStatusEnum{
		INIT: ExtensionVersionSnakeVersionStatus{
			value: "INIT",
		},
		VALIDATING: ExtensionVersionSnakeVersionStatus{
			value: "VALIDATING",
		},
		REJECTED: ExtensionVersionSnakeVersionStatus{
			value: "REJECTED",
		},
		PUBLISHED: ExtensionVersionSnakeVersionStatus{
			value: "PUBLISHED",
		},
		OFFLINE: ExtensionVersionSnakeVersionStatus{
			value: "OFFLINE",
		},
		ABANDONED: ExtensionVersionSnakeVersionStatus{
			value: "ABANDONED",
		},
		GRAY_INIT: ExtensionVersionSnakeVersionStatus{
			value: "GRAY_INIT",
		},
		GRAYED: ExtensionVersionSnakeVersionStatus{
			value: "GRAYED",
		},
		GRAY_REJECTED: ExtensionVersionSnakeVersionStatus{
			value: "GRAY_REJECTED",
		},
	}
}

func (c ExtensionVersionSnakeVersionStatus) Value() string {
	return c.value
}

func (c ExtensionVersionSnakeVersionStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExtensionVersionSnakeVersionStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type ExtensionVersionSnakeVersionValidateStatus struct {
	value string
}

type ExtensionVersionSnakeVersionValidateStatusEnum struct {
	NONE           ExtensionVersionSnakeVersionValidateStatus
	UPLOADING      ExtensionVersionSnakeVersionValidateStatus
	VALIDATING     ExtensionVersionSnakeVersionValidateStatus
	OFFLINING      ExtensionVersionSnakeVersionValidateStatus
	ONLINING       ExtensionVersionSnakeVersionValidateStatus
	UMS_VALIDATING ExtensionVersionSnakeVersionValidateStatus
}

func GetExtensionVersionSnakeVersionValidateStatusEnum() ExtensionVersionSnakeVersionValidateStatusEnum {
	return ExtensionVersionSnakeVersionValidateStatusEnum{
		NONE: ExtensionVersionSnakeVersionValidateStatus{
			value: "NONE",
		},
		UPLOADING: ExtensionVersionSnakeVersionValidateStatus{
			value: "UPLOADING",
		},
		VALIDATING: ExtensionVersionSnakeVersionValidateStatus{
			value: "VALIDATING",
		},
		OFFLINING: ExtensionVersionSnakeVersionValidateStatus{
			value: "OFFLINING",
		},
		ONLINING: ExtensionVersionSnakeVersionValidateStatus{
			value: "ONLINING",
		},
		UMS_VALIDATING: ExtensionVersionSnakeVersionValidateStatus{
			value: "UMS_VALIDATING",
		},
	}
}

func (c ExtensionVersionSnakeVersionValidateStatus) Value() string {
	return c.value
}

func (c ExtensionVersionSnakeVersionValidateStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExtensionVersionSnakeVersionValidateStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
