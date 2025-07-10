package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateDesktopPoolReq 创建桌面请求。
type CreateDesktopPoolReq struct {

	// 桌面池名称，桌面池名称必须保证唯一。桌面名称只允许输入中文、大写字母、小写字母、数字、中划线，长度范围为1~255。
	Name string `json:"name"`

	// 桌面池类型，DYNAMIC：动态池，STATIC：静态池。
	Type CreateDesktopPoolReqType `json:"type"`

	// 桌面池大小：当前桌面池要创建多少台桌面。
	Size int32 `json:"size"`

	// 桌面池描述。
	Description *string `json:"description,omitempty"`

	// 可用分区。将桌面创建到指定的可用分区。如果不指定则使用系统随机的可用分区。
	AvailabilityZone *string `json:"availability_zone,omitempty"`

	// 套餐ID。
	ProductId string `json:"product_id"`

	// 产品规格ID。可用区是边缘可用区时，必填此参数。
	FlavorId *string `json:"flavor_id,omitempty"`

	// 镜像类型。默认值为private。  - private：私有镜像。 - gold：公共镜像。
	ImageType string `json:"image_type"`

	// 镜像ID，用于私有镜像创建桌面场景，配合product_id使用。
	ImageId string `json:"image_id"`

	RootVolume *VolumeInfo `json:"root_volume"`

	// 数据盘列表。
	DataVolumes *[]VolumeInfo `json:"data_volumes,omitempty"`

	// 创建桌面时的VPC ID。
	VpcId *string `json:"vpc_id,omitempty"`

	// 创建桌面使用的子网ID。
	SubnetIds []string `json:"subnet_ids"`

	// 桌面使用的安全组，如果不指定则默认使用桌面代理中指定的安全组。
	SecurityGroups *[]SecurityGroup `json:"security_groups,omitempty"`

	// 要授权的用户/用户组列表。
	AuthorizedObjects *[]AuthorizedObjects `json:"authorized_objects,omitempty"`

	// OU名称，在对接AD时使用，需提前在AD中创建OU。
	OuName *string `json:"ou_name,omitempty"`

	// 标签列表。
	Tags *[]Tag `json:"tags,omitempty"`

	// 企业项目ID，默认\"0。\"
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 动态池桌面断连多少分钟内，保留用户与桌面的绑定关系，超时后自动解绑。
	DisconnectedRetentionPeriod *int32 `json:"disconnected_retention_period,omitempty"`

	// 桌面池是否开启弹性伸缩类型，为false则表示不开启弹性伸缩，为true则表示开启弹性伸缩。
	EnableAutoscale *bool `json:"enable_autoscale,omitempty"`

	AutoscalePolicy *AutoscalePolicy `json:"autoscale_policy,omitempty"`

	// 策略id，用于指定生成桌面名称策略。
	DesktopNamePolicyId *string `json:"desktop_name_policy_id,omitempty"`
}

func (o CreateDesktopPoolReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDesktopPoolReq struct{}"
	}

	return strings.Join([]string{"CreateDesktopPoolReq", string(data)}, " ")
}

type CreateDesktopPoolReqType struct {
	value string
}

type CreateDesktopPoolReqTypeEnum struct {
	DYNAMIC CreateDesktopPoolReqType
	STATIC  CreateDesktopPoolReqType
}

func GetCreateDesktopPoolReqTypeEnum() CreateDesktopPoolReqTypeEnum {
	return CreateDesktopPoolReqTypeEnum{
		DYNAMIC: CreateDesktopPoolReqType{
			value: "DYNAMIC",
		},
		STATIC: CreateDesktopPoolReqType{
			value: "STATIC",
		},
	}
}

func (c CreateDesktopPoolReqType) Value() string {
	return c.value
}

func (c CreateDesktopPoolReqType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateDesktopPoolReqType) UnmarshalJSON(b []byte) error {
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
