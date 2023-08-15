package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateProtectedInstanceRequestParams 创建保护实例数据结构
type CreateProtectedInstanceRequestParams struct {

	// 需要加入的保护组ID。
	ServerGroupId string `json:"server_group_id"`

	// 指定的生产站点云服务器ID。
	ServerId string `json:"server_id"`

	// 指定保护实例的名称，最大支持长度为64个字节。只包含中文字符、英文字母（a～ｚ、Ａ～Ｚ）、数字（０~９）、小数点（．）、下划线（_）、中划线（-）。
	Name string `json:"name"`

	// 指定保护实例的描述，最大支持长度为64个字节。不能包含左尖括号（<）或右尖括号（>）。
	Description *string `json:"description,omitempty"`

	// 专属分布式存储池ID。 当容灾站点磁盘选择专属分布式存储时指定该字段。
	ClusterId *string `json:"cluster_id,omitempty"`

	// 容灾站点云服务器主网卡所在的子网subnetID，与neutron_network_id字段值一致。
	PrimarySubnetId *string `json:"primary_subnet_id,omitempty"`

	// 容灾站点云服务器主网卡IP地址。此参数仅在传入primary_subnet_id时有效。指定primary_subnet_id时，如果不指定该参数，将自动分配容灾站点云服务器主网卡IP地址。
	PrimaryIpAddress *string `json:"primary_ip_address,omitempty"`

	// 标签列表。
	Tags *[]ResourceTag `json:"tags,omitempty"`

	// 指定的容灾站点云服务器的flavor ID。 查询flavor列表，请参见查询云服务器规格变更支持列表。  说明:不指定此参数时，容灾站点云服务器的flavor ID默认和生产站点云服务器保持一致。 不同规格的云服务器在性能上存在差异，可能会对云服务器上运行的应用产生影响。为保证切换/故障切换后云服务器的性能，建议容灾站点服务器的规格（CPU、内存）不低于生产站点云服务器的规格（CPU、内存）。 生产站点云服务器和容灾站点云服务器的flavor存在匹配关系，可以通过上述接口使用生产站点云服务器过滤出满足要求的容灾站点云服务器flavor。
	FlavorRef *string `json:"flavorRef,omitempty"`

	// 在专属主机或共享池中创建容灾站点云服务器，默认为在共享池中创建。 值为：shared或dedicated。shared：表示共享池。 dedicated：表示专属主机。
	Tenancy *CreateProtectedInstanceRequestParamsTenancy `json:"tenancy,omitempty"`

	// 专属主机id，此属性仅在tenancy值为dedicated时有效。 若不指定此属性，系统将自动分配租户可以自动放置弹性云服务器的专属主机。
	DedicatedHostId *string `json:"dedicated_host_id,omitempty"`
}

func (o CreateProtectedInstanceRequestParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateProtectedInstanceRequestParams struct{}"
	}

	return strings.Join([]string{"CreateProtectedInstanceRequestParams", string(data)}, " ")
}

type CreateProtectedInstanceRequestParamsTenancy struct {
	value string
}

type CreateProtectedInstanceRequestParamsTenancyEnum struct {
	SHARED    CreateProtectedInstanceRequestParamsTenancy
	DEDICATED CreateProtectedInstanceRequestParamsTenancy
}

func GetCreateProtectedInstanceRequestParamsTenancyEnum() CreateProtectedInstanceRequestParamsTenancyEnum {
	return CreateProtectedInstanceRequestParamsTenancyEnum{
		SHARED: CreateProtectedInstanceRequestParamsTenancy{
			value: "shared",
		},
		DEDICATED: CreateProtectedInstanceRequestParamsTenancy{
			value: "dedicated",
		},
	}
}

func (c CreateProtectedInstanceRequestParamsTenancy) Value() string {
	return c.value
}

func (c CreateProtectedInstanceRequestParamsTenancy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateProtectedInstanceRequestParamsTenancy) UnmarshalJSON(b []byte) error {
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
