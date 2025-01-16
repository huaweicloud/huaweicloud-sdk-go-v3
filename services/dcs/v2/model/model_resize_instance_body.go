package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ResizeInstanceBody struct {

	// 产品规格编码。具体查询方法如下：  - 方法一：查询产品介绍中的[实例规格](https://support.huaweicloud.com/productdesc-dcs/dcs-pd-0522002.html) - 方法二：登录分布式缓存的控制台界面，点击购买缓存实例，查找对应的实例规格名称 - 方法三：调用[查询产品规格](https://support.huaweicloud.com/api-dcs/ListFlavors.html)接口查询。
	SpecCode string `json:"spec_code"`

	// 新的缓存实例规格，单位：GB。 Redis4.0和Redis5.0：单机和主备类型实例取值：0.125、0.25、0.5、1、2、4、8、16、32、64。Cluster集群实例规格支持24、32、48、64、96、128、192、256、384、512、768、1024。 Memcached：单机和主备类型实例取值：2、4、8、16、32、64。
	NewCapacity int32 `json:"new_capacity"`

	BssParam *BssParamEntity `json:"bss_param,omitempty"`

	// 需要保留的节点ip。cluster集群缩容时需要填写，不填写时系统将随机删除多余的分片
	ReservedIp *[]string `json:"reserved_ip,omitempty"`

	// 变更类型，Redis 4.0或者5.0实例进行副本数变更时必选。 - createReplication: 添加副本 - deleteReplication: 删除副本 - addSharding: 添加分片 - instanceType: 实例类型变更[，目前仅支持Redis 4.0/5.0/6.0实例中的主备实例/读写分离实例变更为proxy集群实例、proxy集群实例变更为主备实例/读写分离实例](tag:hws,hws_hk,hws_eu,ctc,sbc,hk_sbc,g42,hk_g42,otc)。
	ChangeType *ResizeInstanceBodyChangeType `json:"change_type,omitempty"`

	// Redis 4.0或者5.0主备实例进行添加副本时必选，指定每个副本所在的可用区Code，使用前需要先确认该可用区资源是否售罄。  具体查询方法，请参考[查询可用区信息](https://support.huaweicloud.com/api-dcs/ListAvailableZones.html)
	AvailableZones *[]string `json:"available_zones,omitempty"`

	// Redis 4.0或者5.0主备实例进行删除副本时必选，指定需要删除的节点ID，目前仅支持一次删除一个副本。  节点ID查询方法，请参考[查询分片信息](https://support.huaweicloud.com/api-dcs/ListGroupReplicationInfo.html)
	NodeList *[]string `json:"node_list,omitempty"`

	// 是否立即变更。默认值为true。 - true: 立即变更 - false: 可维护时间窗内进行变更
	ExecuteImmediately *bool `json:"execute_immediately,omitempty"`
}

func (o ResizeInstanceBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeInstanceBody struct{}"
	}

	return strings.Join([]string{"ResizeInstanceBody", string(data)}, " ")
}

type ResizeInstanceBodyChangeType struct {
	value string
}

type ResizeInstanceBodyChangeTypeEnum struct {
	CREATE_REPLICATION ResizeInstanceBodyChangeType
	DELETE_REPLICATION ResizeInstanceBodyChangeType
	ADD_SHARDING       ResizeInstanceBodyChangeType
	INSTANCE_TYPE      ResizeInstanceBodyChangeType
}

func GetResizeInstanceBodyChangeTypeEnum() ResizeInstanceBodyChangeTypeEnum {
	return ResizeInstanceBodyChangeTypeEnum{
		CREATE_REPLICATION: ResizeInstanceBodyChangeType{
			value: "createReplication",
		},
		DELETE_REPLICATION: ResizeInstanceBodyChangeType{
			value: "deleteReplication",
		},
		ADD_SHARDING: ResizeInstanceBodyChangeType{
			value: "addSharding",
		},
		INSTANCE_TYPE: ResizeInstanceBodyChangeType{
			value: "instanceType",
		},
	}
}

func (c ResizeInstanceBodyChangeType) Value() string {
	return c.value
}

func (c ResizeInstanceBodyChangeType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ResizeInstanceBodyChangeType) UnmarshalJSON(b []byte) error {
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
