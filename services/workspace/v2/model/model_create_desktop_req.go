package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateDesktopReq 创建桌面请求。
type CreateDesktopReq struct {

	// 云桌面类型。 - DEDICATED：专属桌面，单用户。 - SHARED: 多用户共享桌面。
	DesktopType CreateDesktopReqDesktopType `json:"desktop_type"`

	// 可用分区。将桌面创建到指定的可用分区。如果不指定则使用系统随机的可用分区。
	AvailabilityZone *string `json:"availability_zone,omitempty"`

	// 套餐ID。
	ProductId string `json:"product_id"`

	// 套餐flavorID。
	FlavorId *string `json:"flavor_id,omitempty"`

	// 桌面协同资源SKU码
	ShareResourceSku *string `json:"share_resource_sku,omitempty"`

	// 镜像类型。默认值为private。  - private：私有镜像。 - gold：公共镜像。
	ImageType string `json:"image_type"`

	// 镜像ID，用于私有镜像创建桌面场景，配合product_id使用。
	ImageId string `json:"image_id"`

	RootVolume *Volume `json:"root_volume"`

	// 数据盘列表。
	DataVolumes *[]Volume `json:"data_volumes,omitempty"`

	// 桌面对应的网卡信息，如果不指定则使用默认网卡。
	Nics *[]Nic `json:"nics,omitempty"`

	// 桌面使用的安全组，如果不指定则默认使用桌面代理中指定的安全组。
	SecurityGroups *[]SecurityGroup `json:"security_groups,omitempty"`

	// 创建桌面使用的参数列表。长度为1-100。  当前不支持一批桌面不同配置，所有桌面的配置和第一台的一致，如果第一台未设置参数，则取外层的同名参数。
	Desktops *[]Desktop `json:"desktops,omitempty"`

	// 搭配size使用，当size为1时代表桌面名，位数1-15，当size大于1时代表桌面名前缀，位数：1-13。
	DesktopName *string `json:"desktop_name,omitempty"`

	// 桌面指定分配的ip地址列表,最小为0，最大为100。
	DesktopIps *[]string `json:"desktop_ips,omitempty"`

	// 创建不分配用户的桌面的个数，和desktops不能同时生效，搭配desktop_name使用。
	Size *int32 `json:"size,omitempty"`

	// 创建成功后是否发送邮件通知桌面用户，默认为true。
	EmailNotification *bool `json:"email_notification,omitempty"`

	// 企业ID，在非对接AD场景首次创建桌面时使用。
	EnterpriseId *string `json:"enterprise_id,omitempty"`

	// 企业项目ID，默认\"0\"
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 包周期订购ID，CBC订购回调时使用。
	OrderId *string `json:"order_id,omitempty"`

	// OU名称，在对接AD时使用，需提前在AD中创建OU。
	OuName *string `json:"ou_name,omitempty"`

	// 在非对接AD场景首次创建桌面时使用。
	VpcId *string `json:"vpc_id,omitempty"`

	// 在非对接AD场景首次创建桌面时使用。
	SubnetIds *[]string `json:"subnet_ids,omitempty"`

	// 标签列表。
	Tags *[]Tag `json:"tags,omitempty"`

	SchedulerHints *WdhParam `json:"scheduler_hints,omitempty"`

	// 桌面来源。  - DEFAULT：默认桌面来源 - ONEMOBILE：协同办公云桌面OneMobile
	DesktopIsv *CreateDesktopReqDesktopIsv `json:"desktop_isv,omitempty"`

	// 接入模式。在非对接AD场景首次创建桌面时使用。 - INTERNET：互联网接入。 - DEDICATED：专线接入。 - BOTH：代表两种接入方式都支持。
	AccessMode *CreateDesktopReqAccessMode `json:"access_mode,omitempty"`

	ApplySharedVpcDedicatedParam *ApplySharedVpcDedicatedParam `json:"apply_shared_vpc_dedicated_param,omitempty"`

	// 专线接入网段列表，多个网段信息用分号隔开，列表长度不超过5。在非对接AD场景首次创建桌面时使用。
	DedicatedSubnets *string `json:"dedicated_subnets,omitempty"`

	Eip *Eip `json:"eip,omitempty"`

	Adn *Adn `json:"adn,omitempty"`

	// 专享主机ID，创建专享桌面时如果在指定专享主机中创建则必选
	ExclusiveHostId *string `json:"exclusive_host_id,omitempty"`

	// 策略id，用于指定生成桌面名称策略，如果指定了桌面名称则优先使用指定的桌面名称。
	DesktopNamePolicyId *string `json:"desktop_name_policy_id,omitempty"`

	// 桌面小时包套餐ID。
	HourPackageProductId *string `json:"hour_package_product_id,omitempty"`

	// 桌面小时包offeringID。
	HourPackageOfferingId *string `json:"hour_package_offering_id,omitempty"`

	// 根资源ID列表，创建小时包桌面时使用，最小为0，最大为100。
	RootResourceIds *[]string `json:"root_resource_ids,omitempty"`

	// instInfoId列表，创建小时包桌面时使用，最小为0，最大为100。
	InstInfoIds *[]string `json:"inst_info_ids,omitempty"`
}

func (o CreateDesktopReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDesktopReq struct{}"
	}

	return strings.Join([]string{"CreateDesktopReq", string(data)}, " ")
}

type CreateDesktopReqDesktopType struct {
	value string
}

type CreateDesktopReqDesktopTypeEnum struct {
	DEDICATED CreateDesktopReqDesktopType
	SHARED    CreateDesktopReqDesktopType
}

func GetCreateDesktopReqDesktopTypeEnum() CreateDesktopReqDesktopTypeEnum {
	return CreateDesktopReqDesktopTypeEnum{
		DEDICATED: CreateDesktopReqDesktopType{
			value: "DEDICATED",
		},
		SHARED: CreateDesktopReqDesktopType{
			value: "SHARED",
		},
	}
}

func (c CreateDesktopReqDesktopType) Value() string {
	return c.value
}

func (c CreateDesktopReqDesktopType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateDesktopReqDesktopType) UnmarshalJSON(b []byte) error {
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

type CreateDesktopReqDesktopIsv struct {
	value string
}

type CreateDesktopReqDesktopIsvEnum struct {
	DEFAULT   CreateDesktopReqDesktopIsv
	ONEMOBILE CreateDesktopReqDesktopIsv
}

func GetCreateDesktopReqDesktopIsvEnum() CreateDesktopReqDesktopIsvEnum {
	return CreateDesktopReqDesktopIsvEnum{
		DEFAULT: CreateDesktopReqDesktopIsv{
			value: "DEFAULT",
		},
		ONEMOBILE: CreateDesktopReqDesktopIsv{
			value: "ONEMOBILE",
		},
	}
}

func (c CreateDesktopReqDesktopIsv) Value() string {
	return c.value
}

func (c CreateDesktopReqDesktopIsv) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateDesktopReqDesktopIsv) UnmarshalJSON(b []byte) error {
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

type CreateDesktopReqAccessMode struct {
	value string
}

type CreateDesktopReqAccessModeEnum struct {
	INTERNET  CreateDesktopReqAccessMode
	DEDICATED CreateDesktopReqAccessMode
	BOTH      CreateDesktopReqAccessMode
}

func GetCreateDesktopReqAccessModeEnum() CreateDesktopReqAccessModeEnum {
	return CreateDesktopReqAccessModeEnum{
		INTERNET: CreateDesktopReqAccessMode{
			value: "INTERNET",
		},
		DEDICATED: CreateDesktopReqAccessMode{
			value: "DEDICATED",
		},
		BOTH: CreateDesktopReqAccessMode{
			value: "BOTH",
		},
	}
}

func (c CreateDesktopReqAccessMode) Value() string {
	return c.value
}

func (c CreateDesktopReqAccessMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateDesktopReqAccessMode) UnmarshalJSON(b []byte) error {
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
