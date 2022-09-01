package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// Response Object
type ShowShareResponse struct {
	ActionProgress *ActionProgress `json:"action_progress,omitempty" xml:"action_progress"`

	// SFS Turbo文件系统的版本号。
	Version *string `json:"version,omitempty" xml:"version"`

	// SFS Turbo文件系统剩余容量，单位GB。
	AvailCapacity *string `json:"avail_capacity,omitempty" xml:"avail_capacity"`

	// SFS Turbo文件系统所在可用区编码。
	AvailabilityZone *string `json:"availability_zone,omitempty" xml:"availability_zone"`

	// SFS Turbo文件系统所在可用区名称。
	AzName *string `json:"az_name,omitempty" xml:"az_name"`

	// 创建时间。UTC时间，例如：2018-11-19T04:02:03
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty" xml:"created_at"`

	// 用户指定的加密密钥ID，非加密盘时不返回。
	CryptKeyId *string `json:"crypt_key_id,omitempty" xml:"crypt_key_id"`

	// 如果是增强型文件系统，该字段返回bandwidth，否则不返回。
	ExpandType *string `json:"expand_type,omitempty" xml:"expand_type"`

	// SFS Turbo文件系统的挂载端点。
	ExportLocation *string `json:"export_location,omitempty" xml:"export_location"`

	// SFS Turbo的文件系统ID。
	Id *string `json:"id,omitempty" xml:"id"`

	// 创建时指定的SFS Turbo文件系统名称。
	Name *string `json:"name,omitempty" xml:"name"`

	// SFS Turbo文件系统的计费模式。'0'代表按需付费，'1'代表包周期计费。
	PayModel *ShowShareResponsePayModel `json:"pay_model,omitempty" xml:"pay_model"`

	// SFS Turbo文件系统所在区域。
	Region *string `json:"region,omitempty" xml:"region"`

	// 用户指定的安全组ID。
	SecurityGroupId *string `json:"security_group_id,omitempty" xml:"security_group_id"`

	// SFS Turbo文件系统的协议类型，当前为NFS
	ShareProto *string `json:"share_proto,omitempty" xml:"share_proto"`

	// SFS Turbo文件系统性能类型，包括“STANDARD”标准型和“PERFORMANCE”性能型。
	ShareType *string `json:"share_type,omitempty" xml:"share_type"`

	// SFS Turbo文件系统总容量，单位GB。
	Size *string `json:"size,omitempty" xml:"size"`

	// SFS Turbo文件系统的状态。'100'表示创建中，'200'表示可用，'400'表示已删除，'303'表示创建失败。
	Status *string `json:"status,omitempty" xml:"status"`

	// SFS Turbo文件系统的子状态。'121'表示扩容中，'221'表示扩容成功，'321'表示扩容失败。
	SubStatus *string `json:"sub_status,omitempty" xml:"sub_status"`

	// 用户指定的子网的网络ID。
	SubnetId *string `json:"subnet_id,omitempty" xml:"subnet_id"`

	// 用户指定的VPC ID。
	VpcId          *string `json:"vpc_id,omitempty" xml:"vpc_id"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowShareResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowShareResponse struct{}"
	}

	return strings.Join([]string{"ShowShareResponse", string(data)}, " ")
}

type ShowShareResponsePayModel struct {
	value string
}

type ShowShareResponsePayModelEnum struct {
	E_0 ShowShareResponsePayModel
	E_1 ShowShareResponsePayModel
}

func GetShowShareResponsePayModelEnum() ShowShareResponsePayModelEnum {
	return ShowShareResponsePayModelEnum{
		E_0: ShowShareResponsePayModel{
			value: "0",
		},
		E_1: ShowShareResponsePayModel{
			value: "1",
		},
	}
}

func (c ShowShareResponsePayModel) Value() string {
	return c.value
}

func (c ShowShareResponsePayModel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowShareResponsePayModel) UnmarshalJSON(b []byte) error {
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
