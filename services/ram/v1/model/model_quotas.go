package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Quotas 资源共享的配额。
type Quotas struct {

	// 配额类型。resource_share账号创建资源共享的数量，resource_association资源共享关联的资源数量，principal_association资源共享关联的身份数量，permission_association资源共享关联的权限数量，tag_association资源共享关联的标签数量。
	Type string `json:"type"`

	// 总配额数量。
	Quota int32 `json:"quota"`

	// 最小配额。
	Min int32 `json:"min"`

	// 最大配额。
	Max int32 `json:"max"`

	// 已使用的配额数量。
	Used int32 `json:"used"`
}

func (o Quotas) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Quotas struct{}"
	}

	return strings.Join([]string{"Quotas", string(data)}, " ")
}
