package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

type Shares struct {
	ActionProgress *ActionProgress `json:"action_progress,omitempty"`

	// SFS Turbo文件系统的版本号。
	Version *string `json:"version,omitempty"`

	// SFS Turbo文件系统剩余容量，单位GB。
	AvailCapacity *string `json:"avail_capacity,omitempty"`

	// SFS Turbo文件系统所在可用区编码。
	AvailabilityZone *string `json:"availability_zone,omitempty"`

	// SFS Turbo文件系统所在可用区名称。
	AzName *string `json:"az_name,omitempty"`

	// 创建时间。UTC时间，例如：2018-11-19T04:02:03
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	// 用户指定的加密密钥ID，非加密盘时不返回。
	CryptKeyId *string `json:"crypt_key_id,omitempty"`

	// 如果是增强型文件系统，该字段返回bandwidth，否则不返回。
	ExpandType *string `json:"expand_type,omitempty"`

	// SFS Turbo文件系统的挂载端点。
	ExportLocation *string `json:"export_location,omitempty"`

	// SFS Turbo的文件系统ID。
	Id *string `json:"id,omitempty"`

	// 创建时指定的SFS Turbo文件系统名称。
	Name *string `json:"name,omitempty"`

	// SFS Turbo文件系统的计费模式。'0'代表按需付费，'1'代表包周期计费。
	PayModel *SharesPayModel `json:"pay_model,omitempty"`

	// SFS Turbo文件系统所在区域。
	Region *string `json:"region,omitempty"`

	// 用户指定的安全组ID。
	SecurityGroupId *string `json:"security_group_id,omitempty"`

	// SFS Turbo文件系统的协议类型，当前为NFS
	ShareProto *string `json:"share_proto,omitempty"`

	// SFS Turbo文件系统性能类型，包括“STANDARD”标准型和“PERFORMANCE”性能型。
	ShareType *string `json:"share_type,omitempty"`

	// SFS Turbo文件系统总容量，单位GB。
	Size *string `json:"size,omitempty"`

	// SFS Turbo文件系统的状态。'100'表示创建中，'200'表示可用，'303'表示创建失败。
	Status *string `json:"status,omitempty"`

	// SFS Turbo文件系统的子状态。'121'表示扩容中，'221'表示扩容成功，'321'表示扩容失败。
	SubStatus *string `json:"sub_status,omitempty"`

	// 用户指定的子网的网络ID。
	SubnetId *string `json:"subnet_id,omitempty"`

	// 用户指定的VPC ID。
	VpcId *string `json:"vpc_id,omitempty"`
}

func (o Shares) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Shares struct{}"
	}

	return strings.Join([]string{"Shares", string(data)}, " ")
}

type SharesPayModel struct {
	value string
}

type SharesPayModelEnum struct {
	E_0 SharesPayModel
	E_1 SharesPayModel
}

func GetSharesPayModelEnum() SharesPayModelEnum {
	return SharesPayModelEnum{
		E_0: SharesPayModel{
			value: "0",
		},
		E_1: SharesPayModel{
			value: "1",
		},
	}
}

func (c SharesPayModel) Value() string {
	return c.value
}

func (c SharesPayModel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SharesPayModel) UnmarshalJSON(b []byte) error {
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
