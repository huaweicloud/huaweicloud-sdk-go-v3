package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// Vpc
type Vpc struct {

	// 功能说明：虚拟私有云ID 取值范围：带\"-\"的UUID
	Id string `json:"id"`

	// 功能说明：虚拟私有云名称 取值范围：0-64个字符，支持数字、字母、中文、_(下划线)、-（中划线）、.（点） 约束：如果名称不为空，则同一个租户下的名称不能重复
	Name string `json:"name"`

	// 功能说明：虚拟私有云下可用子网的范围 取值范围： - 10.0.0.0/8~10.255.255.240/28 - 172.16.0.0/12 ~ 172.31.255.240/28 - 192.168.0.0/16 ~ 192.168.255.240/28 不指定cidr时，默认值为空 约束：必须是ipv4 cidr格式，例如:192.168.0.0/16
	Cidr string `json:"cidr"`

	// 功能说明：虚拟私有云的描述 取值范围：0-255个字符，不能包含“<”和“>”
	Description string `json:"description"`

	// 功能说明：路由信息列表，详情参见route对象
	Routes []Route `json:"routes"`

	// 功能说明：虚拟私有云的状态 取值范围： - CREATING：创建中 - OK：创建成功
	Status VpcStatus `json:"status"`

	// 功能说明：企业项目ID。 取值范围：最大长度36字节，带“-”连字符的UUID格式，或者是字符串“0”。“0”表示默认企业项目。
	EnterpriseProjectId string `json:"enterprise_project_id"`

	// 项目ID
	TenantId string `json:"tenant_id"`

	// 功能说明：资源创建UTC时间 格式：yyyy-MM-ddTHH:mm:ss
	CreatedAt *sdktime.SdkTime `json:"created_at"`

	// 功能说明：资源更新UTC时间 格式：yyyy-MM-ddTHH:mm:ss
	UpdatedAt *sdktime.SdkTime `json:"updated_at"`

	// 功能说明：默认情况下，VPC中的资源可以通过内网访问服务终结点。开启该项后，VPC将无法通过内网访问服务终结点，请谨慎操作。 无法访问以下云服务：容器镜像服务SWR、云日志服务LTS、企业主机安全HSS、应用运维管理AOM、应用性能管理APM、对象存储服务OBS、API网关APIG。 取值范围： off：代表禁用。 on：代表开启。
	BlockServiceEndpointStates string `json:"block_service_endpoint_states"`

	// 功能说明：是否开启VPC内所有子网的IPv4地址使用量指标监控。 取值范围： true：开启 false：不开启
	EnableNetworkAddressUsageMetrics bool `json:"enable_network_address_usage_metrics"`
}

func (o Vpc) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Vpc struct{}"
	}

	return strings.Join([]string{"Vpc", string(data)}, " ")
}

type VpcStatus struct {
	value string
}

type VpcStatusEnum struct {
	CREATING VpcStatus
	OK       VpcStatus
	ERROR    VpcStatus
}

func GetVpcStatusEnum() VpcStatusEnum {
	return VpcStatusEnum{
		CREATING: VpcStatus{
			value: "CREATING",
		},
		OK: VpcStatus{
			value: "OK",
		},
		ERROR: VpcStatus{
			value: "ERROR",
		},
	}
}

func (c VpcStatus) Value() string {
	return c.value
}

func (c VpcStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VpcStatus) UnmarshalJSON(b []byte) error {
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
