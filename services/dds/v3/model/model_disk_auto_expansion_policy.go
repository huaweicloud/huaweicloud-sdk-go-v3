package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DiskAutoExpansionPolicy struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 触发自动扩容阈值，只支持输入80、85和90。默认阈值为90，即当已使用存储空间达到总存储空间的90%时就会触发扩容。集群实例的自动扩容阈值指的是每个shard。
	Threshold *int32 `json:"threshold,omitempty"`

	// 扩容步长（s%），默认为10，支持输入10、15和20。当触发自动扩容的时候，自动扩容当前存储空间的s%（非10倍数向上取整。小数点后四舍五入，默认一次最小10G，账户余额不足时，会导致包年包月实例扩容失败）。
	Step *int32 `json:"step,omitempty"`
}

func (o DiskAutoExpansionPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DiskAutoExpansionPolicy struct{}"
	}

	return strings.Join([]string{"DiskAutoExpansionPolicy", string(data)}, " ")
}
