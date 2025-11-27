package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QuotaResource struct {

	// 配额类型，取值范围： - cluster：集群类型配额 - clustergroup：容器舰队类型配额 - rule：规则类型配额 - federation：联邦类型配额
	Type *string `json:"type,omitempty"`

	// 配额的总值
	Quota *string `json:"quota,omitempty"`

	// 已使用配额
	Used *string `json:"used,omitempty"`

	// 单位
	Unit *string `json:"unit,omitempty"`

	// 最小值
	Min *string `json:"min,omitempty"`

	// 最大值
	Max *string `json:"max,omitempty"`
}

func (o QuotaResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotaResource struct{}"
	}

	return strings.Join([]string{"QuotaResource", string(data)}, " ")
}
