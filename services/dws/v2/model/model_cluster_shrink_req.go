package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 集群缩容请求
type ClusterShrinkReq struct {

	// 缩容数
	ShrinkNumber int32 `json:"shrink_number"`

	// 在线
	Online bool `json:"online"`

	// 数据库类型
	Type string `json:"type"`

	// 重试
	Retry *bool `json:"retry,omitempty"`

	// 执行备份
	ForceBackup bool `json:"force_backup"`

	// 是否需要委托
	NeedAgency bool `json:"need_agency"`
}

func (o ClusterShrinkReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterShrinkReq struct{}"
	}

	return strings.Join([]string{"ClusterShrinkReq", string(data)}, " ")
}
