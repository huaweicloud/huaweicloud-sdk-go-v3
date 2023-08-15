package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResizeInstanceReq struct {

	// 规格变更后的规格ID。 若只扩展磁盘大小，则规格ID保持和原实例不变。
	NewSpecCode *string `json:"new_spec_code,omitempty"`

	// 规格变更后的消息存储空间，单位：GB。 若扩展实例基准带宽，则new_storage_space不能低于基准带宽规定的最小磁盘大小。
	NewStorageSpace *int32 `json:"new_storage_space,omitempty"`

	// 扩容类型, 新规格支持扩容类型：\"horizontal\"、\"vertical\"、\"node\"、\"storage\"四种类型。
	OperType *string `json:"oper_type,omitempty"`

	// 扩容后集群节点数。
	NewBrokerNum *int32 `json:"new_broker_num,omitempty"`

	// 新规格变更后的产品ID。 涉及垂直扩容场景，需指定该项。 产品ID可以从[查询产品规格列表](ListEngineProducts.xml)获取。
	NewProductId *string `json:"new_product_id,omitempty"`

	// 实例绑定的弹性IP地址的ID。 以英文逗号隔开多个弹性IP地址的ID。 如果开启了公网再进行扩容，需要填写此参数。
	PublicipId *string `json:"publicip_id,omitempty"`
}

func (o ResizeInstanceReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeInstanceReq struct{}"
	}

	return strings.Join([]string{"ResizeInstanceReq", string(data)}, " ")
}
