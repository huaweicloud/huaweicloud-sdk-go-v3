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
