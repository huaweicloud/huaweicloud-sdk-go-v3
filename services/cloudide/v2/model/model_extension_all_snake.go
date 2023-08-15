package model

import (
	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

type ExtensionAllSnake struct {

	// 插件id
	ExtensionId *string `json:"extension_id,omitempty"`

	// 插件名称
	ExtensionName *string `json:"extension_name,omitempty"`

	// 插件显示名称
	DisplayName *string `json:"display_name,omitempty"`

	// 插件flag;通过传递flag参数来进行过滤或其他操作。flag的基础数字是2\\4\\8\\16;传递的参数只能是这四个数字加法组合而成的数字 利用它们之间二进制的运算获取的值进行其他操作.比如6=0110=0010+0100也就是2和4的集合flags
	Flags *int32 `json:"flags,omitempty"`

	// 更新时间
	LastUpdated *sdktime.SdkTime `json:"last_updated,omitempty"`

	// 上传时间
	PublishedDate *sdktime.SdkTime `json:"published_date,omitempty"`

	// 发布时间
	ReleaseDate *sdktime.SdkTime `json:"release_date,omitempty"`

	// 插件描述
	ShortDescription *string `json:"short_description,omitempty"`

	// 插件标签
	Tags *[]string `json:"tags,omitempty"`

	// 所有标签
	TagAllList *[]string `json:"tag_all_list,omitempty"`

	Publisher *PublisherSnake `json:"publisher,omitempty"`

	// 系统架构
	Arch *[]string `json:"arch,omitempty"`

	// 安装目标
	Target *string `json:"target,omitempty"`

	// 插件分类
	Categories *[]string `json:"categories,omitempty"`

	// 全部分类列表
	CategoryAllList *[]string `json:"category_all_list,omitempty"`

	PublishManager *PublisherSnake `json:"publish_manager,omitempty"`

	// 插件状态  - INIT 上传插件的第一个版本 - NORMAL 插件有审核通过的版本 - OFFLINE 插件下线 - ABANDONED 上传废弃 - GRAYED 灰度插件
	Status *ExtensionAllSnakeStatus `json:"status,omitempty"`

	// 插件审核状态  - NONE 审核结束 - VALIDATING 审核中
	ValidateStatus *ExtensionAllSnakeValidateStatus `json:"validate_status,omitempty"`

	// 下载量
	InstallCount *int32 `json:"install_count,omitempty"`

	// 平均评星值
	AverageStar float32 `json:"average_star,omitempty"`

	// 插件唯一标识内部插件市场保留
	Identifier *string `json:"identifier,omitempty"`

	// 插件支持的操作系统
	SupportOs *[]string `json:"support_os,omitempty"`

	// 插件支持的ide
	SupportIde *int32 `json:"support_ide,omitempty"`

	// 插件支持的ide名称
	SupportIdeInfo *string `json:"support_ide_info,omitempty"`

	// 插件版本集合
	Versions *[]ExtensionVersionSnake `json:"versions,omitempty"`

	// 插件审核结果
	ValidateResult *string `json:"validate_result,omitempty"`

	ExtensionStatistics *ExtensionStatistics `json:"extension_statistics,omitempty"`

	// 是否支持预览
	Preview *bool `json:"preview,omitempty"`

	ExtInfo *ExtensionExternalInfo `json:"ext_info,omitempty"`

	// 安装目标
	Platform *string `json:"platform,omitempty"`

	CheckResult *CheckResult `json:"check_result,omitempty"`

	// 灰度版本数量
	GrayVersionCount *int32 `json:"gray_version_count,omitempty"`

	// 插件作者
	ExtensionOwner *string `json:"extension_owner,omitempty"`
}

func (o ExtensionAllSnake) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExtensionAllSnake struct{}"
	}

	return strings.Join([]string{"ExtensionAllSnake", string(data)}, " ")
}

type ExtensionAllSnakeStatus struct {
	value string
}

type ExtensionAllSnakeStatusEnum struct {
	INIT      ExtensionAllSnakeStatus
	NORMAL    ExtensionAllSnakeStatus
	OFFLINE   ExtensionAllSnakeStatus
	ABANDONED ExtensionAllSnakeStatus
	GRAYED    ExtensionAllSnakeStatus
}

func GetExtensionAllSnakeStatusEnum() ExtensionAllSnakeStatusEnum {
	return ExtensionAllSnakeStatusEnum{
		INIT: ExtensionAllSnakeStatus{
			value: "INIT",
		},
		NORMAL: ExtensionAllSnakeStatus{
			value: "NORMAL",
		},
		OFFLINE: ExtensionAllSnakeStatus{
			value: "OFFLINE",
		},
		ABANDONED: ExtensionAllSnakeStatus{
			value: "ABANDONED",
		},
		GRAYED: ExtensionAllSnakeStatus{
			value: "GRAYED",
		},
	}
}

func (c ExtensionAllSnakeStatus) Value() string {
	return c.value
}

func (c ExtensionAllSnakeStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExtensionAllSnakeStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}

type ExtensionAllSnakeValidateStatus struct {
	value string
}

type ExtensionAllSnakeValidateStatusEnum struct {
	NONE       ExtensionAllSnakeValidateStatus
	VALIDATING ExtensionAllSnakeValidateStatus
}

func GetExtensionAllSnakeValidateStatusEnum() ExtensionAllSnakeValidateStatusEnum {
	return ExtensionAllSnakeValidateStatusEnum{
		NONE: ExtensionAllSnakeValidateStatus{
			value: "NONE",
		},
		VALIDATING: ExtensionAllSnakeValidateStatus{
			value: "VALIDATING",
		},
	}
}

func (c ExtensionAllSnakeValidateStatus) Value() string {
	return c.value
}

func (c ExtensionAllSnakeValidateStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExtensionAllSnakeValidateStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
