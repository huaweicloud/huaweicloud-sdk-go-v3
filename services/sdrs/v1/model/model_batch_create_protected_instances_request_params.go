package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 批量创建保护实例请求数据接口
type BatchCreateProtectedInstancesRequestParams struct {

	// 保护实例的名称前缀，批量创建保护实例时，为区分不同保护实例，创建过程中系统会自动在名称后加\"-0001\"的类似标记，故此时名称的长度为[1-59]个字符。只包含中文字符、英文字母（a～ｚ、Ａ～Ｚ）、数字（０~９）、小数点（．）、下划线（_）、中划线（-）。
	NamePrefix string `json:"name_prefix"`

	// 保护实例的描述，最大支持长度为64个字符。不能包含左尖括号（<）或右尖括号（>）。
	Description *string `json:"description,omitempty"`

	// 需要加入的保护组ID。
	ServerGroupId string `json:"server_group_id"`

	// 专属分布式存储池ID。当容灾站点磁盘选择专属分布式存储时指定该字段。
	ClusterId *string `json:"cluster_id,omitempty"`

	// 容灾站点云服务器主网卡所在的子网subnetID，与neutron_network_id字段值一致。
	PrimarySubnetId *string `json:"primary_subnet_id,omitempty"`

	// 在专属主机或共享池中创建容灾站点云服务器，默认为在共享池中创建。值为：shared或dedicated。shared：表示共享池。dedicated：表示专属主机。
	Tenancy *BatchCreateProtectedInstancesRequestParamsTenancy `json:"tenancy,omitempty"`

	// 专属主机id，此属性仅在tenancy值为dedicated时有效。若不指定此属性，系统将自动分配租户可以自动放置弹性云服务器的专属主机。
	DedicatedHostId *string `json:"dedicated_host_id,omitempty"`

	// 用于创建保护实例的云服务器信息列表。
	Servers []ServerInfo `json:"servers"`

	// 标签列表。
	Tags *[]ResourceTag `json:"tags,omitempty"`
}

func (o BatchCreateProtectedInstancesRequestParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateProtectedInstancesRequestParams struct{}"
	}

	return strings.Join([]string{"BatchCreateProtectedInstancesRequestParams", string(data)}, " ")
}

type BatchCreateProtectedInstancesRequestParamsTenancy struct {
	value string
}

type BatchCreateProtectedInstancesRequestParamsTenancyEnum struct {
	SHARED    BatchCreateProtectedInstancesRequestParamsTenancy
	DEDICATED BatchCreateProtectedInstancesRequestParamsTenancy
}

func GetBatchCreateProtectedInstancesRequestParamsTenancyEnum() BatchCreateProtectedInstancesRequestParamsTenancyEnum {
	return BatchCreateProtectedInstancesRequestParamsTenancyEnum{
		SHARED: BatchCreateProtectedInstancesRequestParamsTenancy{
			value: "shared",
		},
		DEDICATED: BatchCreateProtectedInstancesRequestParamsTenancy{
			value: "dedicated",
		},
	}
}

func (c BatchCreateProtectedInstancesRequestParamsTenancy) Value() string {
	return c.value
}

func (c BatchCreateProtectedInstancesRequestParamsTenancy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchCreateProtectedInstancesRequestParamsTenancy) UnmarshalJSON(b []byte) error {
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
