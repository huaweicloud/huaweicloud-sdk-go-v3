package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClusterBaselineResponseInfo 基线检查项信息
type ClusterBaselineResponseInfo struct {

	// 检查项描述
	BaselineDesc *string `json:"baseline_desc,omitempty"`

	// 检查项ID
	BaselineIndex *string `json:"baseline_index,omitempty"`

	// 检查项类型
	BaselineType *string `json:"baseline_type,omitempty"`
}

func (o ClusterBaselineResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterBaselineResponseInfo struct{}"
	}

	return strings.Join([]string{"ClusterBaselineResponseInfo", string(data)}, " ")
}
