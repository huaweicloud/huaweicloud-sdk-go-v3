package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SecurityTarget struct {

	// 安全态势感知配置绑定的对象，目前仅支持PRODUCT产品级别，仅对设备级别的安全态势感知项生效。
	TargetType *string `json:"target_type,omitempty"`

	// 绑定对象id列表，target_type为PRODUCT时，由于产品ID在不同资源空间下可以重复，target_id格式为：资源空间ID:产品ID；资源空间ID与产品ID使用冒号拼接而成。
	TargetIds *[]string `json:"target_ids,omitempty"`
}

func (o SecurityTarget) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityTarget struct{}"
	}

	return strings.Join([]string{"SecurityTarget", string(data)}, " ")
}
