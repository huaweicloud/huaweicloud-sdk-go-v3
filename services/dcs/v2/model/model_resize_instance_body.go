package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

type ResizeInstanceBody struct {
	// 产品规格编码。具体查询方法如下：  - 方法一：查询产品介绍中的[实例规格](https://support.huaweicloud.com/productdesc-dcs/dcs-pd-0522002.html) - 方法二：登录分布式缓存的控制台界面，点击购买缓存实例，查找对应的实例规格名称 - 方法三：调用[查询产品规格](https://support.huaweicloud.com/api-dcs/ListFlavors.html)接口查询。

	SpecCode string `json:"spec_code"`
	// 新的缓存实例规格，新的规格必须大于扩容前的规格，单位：GB。 取值包括：4，8，16，32，64 取值必须是当前产品支持的实例规格，请以实际为准。

	NewCapacity int32 `json:"new_capacity"`

	BssParam *BssParamEntity `json:"bss_param,omitempty"`
	// 需要保留的节点ip。cluster集群缩容时需要填写，不填写时系统将随机删除多余的分片

	ReservedIp *[]string `json:"reserved_ip,omitempty"`
	// 变更类型，Redis 4.0或者5.0主备实例进行副本数变更时必选。 - createReplication: 添加副本 - deleteReplication: 删除副本

	ChangeType *ResizeInstanceBodyChangeType `json:"change_type,omitempty"`
	// Redis 4.0或者5.0主备实例进行添加副本时必选，指定每个副本所在的可用区Code，使用前需要先确认该可用区资源是否售罄。  具体查询方法，请参考[查询可用区信息](https://support.huaweicloud.com/api-dcs/ListAvailableZones.html)

	AvailableZones *[]string `json:"available_zones,omitempty"`
	// Redis 4.0或者5.0主备实例进行删除副本时必选，指定需要删除的节点ID，目前仅支持一次删除一个副本。  节点ID查询方法，请参考[查询分片信息](https://support.huaweicloud.com/api-dcs/ListGroupReplicationInfo.html)

	NodeList *[]string `json:"node_list,omitempty"`
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
}

func GetResizeInstanceBodyChangeTypeEnum() ResizeInstanceBodyChangeTypeEnum {
	return ResizeInstanceBodyChangeTypeEnum{
		CREATE_REPLICATION: ResizeInstanceBodyChangeType{
			value: "createReplication",
		},
		DELETE_REPLICATION: ResizeInstanceBodyChangeType{
			value: "deleteReplication",
		},
	}
}

func (c ResizeInstanceBodyChangeType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ResizeInstanceBodyChangeType) UnmarshalJSON(b []byte) error {
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
