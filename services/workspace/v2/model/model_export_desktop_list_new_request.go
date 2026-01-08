package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ExportDesktopListNewRequest Request Object
type ExportDesktopListNewRequest struct {

	// 桌面名。
	ComputerName *string `json:"computer_name,omitempty"`

	// 桌面IP地址。
	DesktopIp *string `json:"desktop_ip,omitempty"`

	// 桌面ID。
	DesktopId *string `json:"desktop_id,omitempty"`

	// 桌面的标签。标签的键和标签的值用“=”连接。
	Tag *string `json:"tag,omitempty"`

	// 语言。  - zh_CN：中文 - en_US：英文
	Language ExportDesktopListNewRequestLanguage `json:"language"`

	// 企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 桌面类型，为空时查所有桌面。 - DEDICATED：普通桌面，包括专享桌面、专属桌面等 - POOLED：池桌面，即桌面池里的桌面
	DesktopType *string `json:"desktop_type,omitempty"`

	// 桌面状态。  - ACTIVE：运行中。 - SHUTOFF：关机。 - ERROR：异常。 - HIBERNATED：休眠。
	Status *string `json:"status,omitempty"`

	// 桌面所属用户，批量筛选，最多不超过100个用户。
	UserNames *[]string `json:"user_names,omitempty"`

	// 排序字段名称，需要结合sort_type字段一起使用。 - created 创建时间。 - computer_name 桌面名称。
	SortField *ExportDesktopListNewRequestSortField `json:"sort_field,omitempty"`

	// 排序类型，默认升序，需要结合sort_field字段一起使用。 - ASC 升序。 - DESC 降序。
	SortType *ExportDesktopListNewRequestSortType `json:"sort_type,omitempty"`

	// 桌面池ID,多个桌面池ID用逗号隔开。
	PoolId *string `json:"pool_id,omitempty"`

	// 是否分配了用户。
	UserAttached *bool `json:"user_attached,omitempty"`

	// 镜像ID。
	ImageId *string `json:"image_id,omitempty"`

	// 计费模式，0：包周期，1：按需。
	ChargeMode *string `json:"charge_mode,omitempty"`

	// 按照维护模式过滤。
	InMaintenanceMode *bool `json:"in_maintenance_mode,omitempty"`

	// 桌面的子网ID。
	SubnetId *string `json:"subnet_id,omitempty"`

	// 连接状态版本，默认值为OLD。 - NEW：新版本 - OLD：老版本
	ConnectionStatusVersion *string `json:"connection_status_version,omitempty"`
}

func (o ExportDesktopListNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportDesktopListNewRequest struct{}"
	}

	return strings.Join([]string{"ExportDesktopListNewRequest", string(data)}, " ")
}

type ExportDesktopListNewRequestLanguage struct {
	value string
}

type ExportDesktopListNewRequestLanguageEnum struct {
	ZH_CN ExportDesktopListNewRequestLanguage
	EN_US ExportDesktopListNewRequestLanguage
}

func GetExportDesktopListNewRequestLanguageEnum() ExportDesktopListNewRequestLanguageEnum {
	return ExportDesktopListNewRequestLanguageEnum{
		ZH_CN: ExportDesktopListNewRequestLanguage{
			value: "zh_CN",
		},
		EN_US: ExportDesktopListNewRequestLanguage{
			value: "en_US",
		},
	}
}

func (c ExportDesktopListNewRequestLanguage) Value() string {
	return c.value
}

func (c ExportDesktopListNewRequestLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExportDesktopListNewRequestLanguage) UnmarshalJSON(b []byte) error {
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

type ExportDesktopListNewRequestSortField struct {
	value string
}

type ExportDesktopListNewRequestSortFieldEnum struct {
	CREATED       ExportDesktopListNewRequestSortField
	COMPUTER_NAME ExportDesktopListNewRequestSortField
}

func GetExportDesktopListNewRequestSortFieldEnum() ExportDesktopListNewRequestSortFieldEnum {
	return ExportDesktopListNewRequestSortFieldEnum{
		CREATED: ExportDesktopListNewRequestSortField{
			value: "created",
		},
		COMPUTER_NAME: ExportDesktopListNewRequestSortField{
			value: "computer_name",
		},
	}
}

func (c ExportDesktopListNewRequestSortField) Value() string {
	return c.value
}

func (c ExportDesktopListNewRequestSortField) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExportDesktopListNewRequestSortField) UnmarshalJSON(b []byte) error {
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

type ExportDesktopListNewRequestSortType struct {
	value string
}

type ExportDesktopListNewRequestSortTypeEnum struct {
	ASC  ExportDesktopListNewRequestSortType
	DESC ExportDesktopListNewRequestSortType
}

func GetExportDesktopListNewRequestSortTypeEnum() ExportDesktopListNewRequestSortTypeEnum {
	return ExportDesktopListNewRequestSortTypeEnum{
		ASC: ExportDesktopListNewRequestSortType{
			value: "ASC",
		},
		DESC: ExportDesktopListNewRequestSortType{
			value: "DESC",
		},
	}
}

func (c ExportDesktopListNewRequestSortType) Value() string {
	return c.value
}

func (c ExportDesktopListNewRequestSortType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExportDesktopListNewRequestSortType) UnmarshalJSON(b []byte) error {
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
