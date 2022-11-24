package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DiskAutoExpansionPolicy struct {

	// 触发自动扩容阈值，只支持输入80、85和90。默认阈值为90，即当已使用存储空间达到总存储空间的90%或者可用空间小于10GB时就会触发扩容。
	Threshold *int32 `json:"threshold,omitempty"`

	// 扩容步长（s%），默认为10，支持输入10、15和20。当触发自动扩容的时候，自动扩容当前存储空间的s%（非10倍数向上取整。小数点后四舍五入，默认一次最小100G，账户余额不足时，会导致包年包月实例扩容失败）
	Step *int32 `json:"step,omitempty"`

	// 实例通过自动扩容所能达到的存储空间上限,单位：GB。需要大于等于实例购买的存储空间大小，且最大上限不能超过实例当前规格支持的最大存储容量。批量自动扩容时，不支持自定义存储自动扩容上限，默认扩容至所选实例对应的最大存储空间。
	Size *int32 `json:"size,omitempty"`
}

func (o DiskAutoExpansionPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DiskAutoExpansionPolicy struct{}"
	}

	return strings.Join([]string{"DiskAutoExpansionPolicy", string(data)}, " ")
}
