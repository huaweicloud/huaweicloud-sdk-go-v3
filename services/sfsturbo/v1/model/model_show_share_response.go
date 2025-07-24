package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// ShowShareResponse Response Object
type ShowShareResponse struct {
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

	// 如果是增强版文件系统，该字段返回bandwidth；如果是20MB/s/TiB、40MB/s/TiB、125MB/s/TiB、250MB/s/TiB、500MB/s/TiB、1000MB/s/TiB，该字段返回hpc；否则不返回。
	ExpandType *string `json:"expand_type,omitempty"`

	// SFS Turbo文件系统的挂载端点。例如\"192.168.0.90:/\"。如果文件系统正在创建，该字段不返回。
	ExportLocation *string `json:"export_location,omitempty"`

	// SFS Turbo的文件系统ID。
	Id *string `json:"id,omitempty"`

	// 创建时指定的SFS Turbo文件系统名称。
	Name *string `json:"name,omitempty"`

	// SFS Turbo文件系统的计费模式。'0'代表按需付费，'1'代表包周期计费。如果文件系统正在创建，该字段不返回。
	PayModel *ShowShareResponsePayModel `json:"pay_model,omitempty"`

	// SFS Turbo文件系统所在区域。
	Region *string `json:"region,omitempty"`

	// 用户指定的安全组ID。
	SecurityGroupId *string `json:"security_group_id,omitempty"`

	// SFS Turbo文件系统的协议类型，当前为NFS或 CIFS。
	ShareProto *string `json:"share_proto,omitempty"`

	// SFS Turbo文件系统性能类型，包括“STANDARD”标准型和“PERFORMANCE”性能型。
	ShareType *string `json:"share_type,omitempty"`

	// SFS Turbo文件系统总容量，单位GB。
	Size *string `json:"size,omitempty"`

	// SFS Turbo文件系统的状态。'100'表示创建中，'200'表示可用，'303'表示创建失败，'800'表示实例被冻结。
	Status *string `json:"status,omitempty"`

	// SFS Turbo文件系统的子状态。当用户未对文件系统有修改类操作时，该字段不返回。 '121'表示扩容中；'132'表示修改安全组中；'137'表示添加VPC中；'138'表示删除VPC中；'150'表示配置联动后端中；'151'表示删除联动后端配置中。 '221'表示扩容成功；'232'表示修改安全组成功；'237'表示添加VPC成功；'238'表示删除VPC成功；'250'表示配置联动后端成功；'251'表示删除联动后端配置成功。 '321'表示扩容失败；'332'表示修改安全组失败；'337'表示添加VPC失败；'338'表示删除VPC失败；'350'表示配置联动后端失败；'351'表示删除联动后端配置失败。
	SubStatus *string `json:"sub_status,omitempty"`

	// 用户指定的子网的网络ID。
	SubnetId *string `json:"subnet_id,omitempty"`

	// 用户指定的VPC ID。
	VpcId *string `json:"vpc_id,omitempty"`

	// SFS Turbo文件系统绑定的企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// tag标签的列表。
	Tags *[]ResourceTag `json:"tags,omitempty"`

	// 可选的挂载IP地址。上一代文件系统规格类型不返回该字段。
	OptionalEndpoint *string `json:"optional_endpoint,omitempty"`

	// 文件系统的带宽规格。 - \"20M\"表示20MB/s/TiB - \"40M\"表示40MB/s/TiB - \"125M\"表示125MB/s/TiB - \"250M\"表示250MB/s/TiB - \"500M\"表示500MB/s/TiB  - \"1000M\"表示1000MB/s/TiB  - \"2G\"、\"4G\"、\"8G\"、\"16G\"、\"24G\"、\"32G\"或\"48G\"表示HPC缓存型的带宽规格。
	HpcBw *string `json:"hpc_bw,omitempty"`

	// 文件系统规格的节点id，为预留字段，不具备实际含义。
	InstanceId *string `json:"instanceId,omitempty"`

	// 文件系统规格的节点类型，为预留字段，不具备实际含义。
	InstanceType *string `json:"instanceType,omitempty"`

	// 文件系统的请求ID，为预留字段，不具备实际含义。
	StatusDetail *string `json:"statusDetail,omitempty"`

	Features       *ShareInfoFeatures `json:"features,omitempty"`
	HttpStatusCode int                `json:"-"`
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
