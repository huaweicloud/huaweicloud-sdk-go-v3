package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListPoolDesktopsDetailRequest Request Object
type ListPoolDesktopsDetailRequest struct {

	// 桌面池ID。
	PoolId string `json:"pool_id"`

	// 通过该类型过滤出与桌面池规格类型不一致的桌面 - PRODUCT_ID: 查找productID与桌面池套餐ID不一致的桌面 - IMAGE_ID: 查找imageID与桌面池镜像ID不一致的桌面 - DISK_NUM: 查找数据盘数量与桌面池配置不一致的桌面 - DISK_SIZE: 查找磁盘累加容量与桌面池配置不一致的桌面
	InconsistentType *ListPoolDesktopsDetailRequestInconsistentType `json:"inconsistent_type,omitempty"`

	// 用于分页查询，查询的起始记录序号，从0开始。
	Offset *int32 `json:"offset,omitempty"`

	// 用于分页查询，返回桌面数量限制。取值范围0-100，默认值是10。
	Limit *int32 `json:"limit,omitempty"`

	// 桌面状态。  - ACTIVE：运行中。 - SHUTOFF：关机。 - ERROR：异常。 - HIBERNATED：休眠。
	Status *string `json:"status,omitempty"`

	// 桌面所属用户，当传user_names时，本字段不生效。
	UserName *string `json:"user_name,omitempty"`

	// 桌面所属用户，批量筛选，最多不超过100个用户。
	UserNames *[]string `json:"user_names,omitempty"`

	// 排序字段名称，需要结合sort_type字段一起使用。 - created 创建时间。 - computer_name 桌面名称。
	SortField *ListPoolDesktopsDetailRequestSortField `json:"sort_field,omitempty"`

	// 排序类型，默认升序，需要结合sort_field字段一起使用。 - ASC 升序。 - DESC 降序。
	SortType *ListPoolDesktopsDetailRequestSortType `json:"sort_type,omitempty"`

	// 桌面名。
	ComputerName *string `json:"computer_name,omitempty"`

	// 桌面名列表。
	ComputerNames *[]string `json:"computer_names,omitempty"`

	// 桌面IP地址。
	DesktopIp *string `json:"desktop_ip,omitempty"`

	// 桌面ID。
	DesktopId *[]string `json:"desktop_id,omitempty"`

	// 桌面类型，为空时查所有桌面。 - DEDICATED：普通桌面，包括专享桌面、专属桌面等 - POOLED：池桌面，即桌面池里的桌面
	DesktopType *string `json:"desktop_type,omitempty"`

	// 桌面的标签。样例：  - key1=value1。 - key1=value1，key2=value2。
	Tag *string `json:"tag,omitempty"`

	// 是否分配了用户。
	UserAttached *bool `json:"user_attached,omitempty"`

	// 企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 镜像ID。
	ImageId *string `json:"image_id,omitempty"`

	// 计费模式，0：包周期，1：按需。
	ChargeMode *string `json:"charge_mode,omitempty"`

	// 按照维护模式过滤。
	InMaintenanceMode *bool `json:"in_maintenance_mode,omitempty"`

	// 是否为协同桌面。
	IsShareDesktop *bool `json:"is_share_desktop,omitempty"`

	// 桌面的子网ID。
	SubnetId *string `json:"subnet_id,omitempty"`

	// 是否支持上网。
	IsSupportInternet *bool `json:"is_support_internet,omitempty"`

	// 查询可用区。
	AvailabilityZone *string `json:"availability_zone,omitempty"`
}

func (o ListPoolDesktopsDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPoolDesktopsDetailRequest struct{}"
	}

	return strings.Join([]string{"ListPoolDesktopsDetailRequest", string(data)}, " ")
}

type ListPoolDesktopsDetailRequestInconsistentType struct {
	value string
}

type ListPoolDesktopsDetailRequestInconsistentTypeEnum struct {
	PRODUCT_ID ListPoolDesktopsDetailRequestInconsistentType
	IMAGE_ID   ListPoolDesktopsDetailRequestInconsistentType
	DISK_NUM   ListPoolDesktopsDetailRequestInconsistentType
	DISK_SIZE  ListPoolDesktopsDetailRequestInconsistentType
}

func GetListPoolDesktopsDetailRequestInconsistentTypeEnum() ListPoolDesktopsDetailRequestInconsistentTypeEnum {
	return ListPoolDesktopsDetailRequestInconsistentTypeEnum{
		PRODUCT_ID: ListPoolDesktopsDetailRequestInconsistentType{
			value: "PRODUCT_ID",
		},
		IMAGE_ID: ListPoolDesktopsDetailRequestInconsistentType{
			value: "IMAGE_ID",
		},
		DISK_NUM: ListPoolDesktopsDetailRequestInconsistentType{
			value: "DISK_NUM",
		},
		DISK_SIZE: ListPoolDesktopsDetailRequestInconsistentType{
			value: "DISK_SIZE",
		},
	}
}

func (c ListPoolDesktopsDetailRequestInconsistentType) Value() string {
	return c.value
}

func (c ListPoolDesktopsDetailRequestInconsistentType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListPoolDesktopsDetailRequestInconsistentType) UnmarshalJSON(b []byte) error {
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

type ListPoolDesktopsDetailRequestSortField struct {
	value string
}

type ListPoolDesktopsDetailRequestSortFieldEnum struct {
	CREATED       ListPoolDesktopsDetailRequestSortField
	COMPUTER_NAME ListPoolDesktopsDetailRequestSortField
}

func GetListPoolDesktopsDetailRequestSortFieldEnum() ListPoolDesktopsDetailRequestSortFieldEnum {
	return ListPoolDesktopsDetailRequestSortFieldEnum{
		CREATED: ListPoolDesktopsDetailRequestSortField{
			value: "created",
		},
		COMPUTER_NAME: ListPoolDesktopsDetailRequestSortField{
			value: "computer_name",
		},
	}
}

func (c ListPoolDesktopsDetailRequestSortField) Value() string {
	return c.value
}

func (c ListPoolDesktopsDetailRequestSortField) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListPoolDesktopsDetailRequestSortField) UnmarshalJSON(b []byte) error {
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

type ListPoolDesktopsDetailRequestSortType struct {
	value string
}

type ListPoolDesktopsDetailRequestSortTypeEnum struct {
	ASC  ListPoolDesktopsDetailRequestSortType
	DESC ListPoolDesktopsDetailRequestSortType
}

func GetListPoolDesktopsDetailRequestSortTypeEnum() ListPoolDesktopsDetailRequestSortTypeEnum {
	return ListPoolDesktopsDetailRequestSortTypeEnum{
		ASC: ListPoolDesktopsDetailRequestSortType{
			value: "ASC",
		},
		DESC: ListPoolDesktopsDetailRequestSortType{
			value: "DESC",
		},
	}
}

func (c ListPoolDesktopsDetailRequestSortType) Value() string {
	return c.value
}

func (c ListPoolDesktopsDetailRequestSortType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListPoolDesktopsDetailRequestSortType) UnmarshalJSON(b []byte) error {
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
