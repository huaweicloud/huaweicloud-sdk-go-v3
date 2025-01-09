package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListDesktopsDetailRequest Request Object
type ListDesktopsDetailRequest struct {

	// 桌面状态。  - ACTIVE：运行中。 - SHUTOFF：关机。 - ERROR：异常。 - HIBERNATED：休眠。
	Status *string `json:"status,omitempty"`

	// 桌面所属用户，当传user_names时，本字段不生效。
	UserName *string `json:"user_name,omitempty"`

	// 桌面所属用户，批量筛选，最多不超过100个用户。
	UserNames *[]string `json:"user_names,omitempty"`

	// 排序字段名称，需要结合sort_type字段一起使用。 - created 创建时间。 - computer_name 桌面名称。
	SortField *ListDesktopsDetailRequestSortField `json:"sort_field,omitempty"`

	// 排序类型，默认升序，需要结合sort_field字段一起使用。 - ASC 升序。 - DESC 降序。
	SortType *ListDesktopsDetailRequestSortType `json:"sort_type,omitempty"`

	// 桌面名。
	ComputerName *string `json:"computer_name,omitempty"`

	// 桌面IP地址。
	DesktopIp *string `json:"desktop_ip,omitempty"`

	// 桌面的sid列表。
	Sids *string `json:"sids,omitempty"`

	// 用于分页查询，查询的起始记录序号，从0开始。
	Offset *int32 `json:"offset,omitempty"`

	// 用于分页查询，取值范围0-500，默认值500。
	Limit *int32 `json:"limit,omitempty"`

	// 桌面ID。
	DesktopId *[]string `json:"desktop_id,omitempty"`

	// 桌面类型，为空时查所有桌面。 - DEDICATED：普通桌面，包括专享桌面、专属桌面等 - POOLED：池桌面，即桌面池里的桌面
	DesktopType *string `json:"desktop_type,omitempty"`

	// 桌面的登录状态。  - UNREGISTER：表示桌面未注册时的状态（桌面启动后，会自动注册）。关机后也会出现未注册的状态。 - REGISTERED：表示桌面注册以后，等待用户连接的状态。 - CONNECTED：表示用户已经成功登录，正在使用桌面。 - DISCONNECTED：表示桌面与客户端断开会话后显示的状态，可能为关闭客户端窗口，或客户端与桌面网络断开引起。
	LoginStatus *string `json:"login_status,omitempty"`

	// 桌面的标签。样例：  - key1=value1。 - key1=value1，key2=value2。
	Tag *string `json:"tag,omitempty"`

	// 是不是模糊用户查询。
	FuzzyQuery *bool `json:"fuzzy_query,omitempty"`

	// 桌面池ID,多个桌面池ID用逗号隔开。
	PoolId *string `json:"pool_id,omitempty"`

	// 是否分配了用户。
	UserAttached *bool `json:"user_attached,omitempty"`

	// 企业项目ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 镜像ID
	ImageId *string `json:"image_id,omitempty"`

	// 计费模式，0：包周期，1：按需。
	ChargeMode *string `json:"charge_mode,omitempty"`

	// 按照维护模式过滤
	InMaintenanceMode *bool `json:"in_maintenance_mode,omitempty"`

	// 是否为协同桌面
	IsShareDesktop *bool `json:"is_share_desktop,omitempty"`

	// 桌面的子网ID。
	SubnetId *string `json:"subnet_id,omitempty"`

	// 是否支持上网。
	IsSupportInternet *bool `json:"is_support_internet,omitempty"`

	// 查询可用区
	AvailabilityZone *string `json:"availability_zone,omitempty"`
}

func (o ListDesktopsDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDesktopsDetailRequest struct{}"
	}

	return strings.Join([]string{"ListDesktopsDetailRequest", string(data)}, " ")
}

type ListDesktopsDetailRequestSortField struct {
	value string
}

type ListDesktopsDetailRequestSortFieldEnum struct {
	CREATED       ListDesktopsDetailRequestSortField
	COMPUTER_NAME ListDesktopsDetailRequestSortField
}

func GetListDesktopsDetailRequestSortFieldEnum() ListDesktopsDetailRequestSortFieldEnum {
	return ListDesktopsDetailRequestSortFieldEnum{
		CREATED: ListDesktopsDetailRequestSortField{
			value: "created",
		},
		COMPUTER_NAME: ListDesktopsDetailRequestSortField{
			value: "computer_name",
		},
	}
}

func (c ListDesktopsDetailRequestSortField) Value() string {
	return c.value
}

func (c ListDesktopsDetailRequestSortField) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListDesktopsDetailRequestSortField) UnmarshalJSON(b []byte) error {
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

type ListDesktopsDetailRequestSortType struct {
	value string
}

type ListDesktopsDetailRequestSortTypeEnum struct {
	ASC  ListDesktopsDetailRequestSortType
	DESC ListDesktopsDetailRequestSortType
}

func GetListDesktopsDetailRequestSortTypeEnum() ListDesktopsDetailRequestSortTypeEnum {
	return ListDesktopsDetailRequestSortTypeEnum{
		ASC: ListDesktopsDetailRequestSortType{
			value: "ASC",
		},
		DESC: ListDesktopsDetailRequestSortType{
			value: "DESC",
		},
	}
}

func (c ListDesktopsDetailRequestSortType) Value() string {
	return c.value
}

func (c ListDesktopsDetailRequestSortType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListDesktopsDetailRequestSortType) UnmarshalJSON(b []byte) error {
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
