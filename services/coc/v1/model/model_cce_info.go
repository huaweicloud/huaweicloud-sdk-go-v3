package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CceInfo cce集群信息
type CceInfo struct {

	// 主键id
	Id *string `json:"id,omitempty"`

	// cce集群id
	CceId *string `json:"cce_id,omitempty"`

	// cce集群名称
	CceName *string `json:"cce_name,omitempty"`

	// 合规数量
	CompliantCount *string `json:"compliant_count,omitempty"`

	// 不合规数量
	NonCompliantCount *string `json:"non_compliant_count,omitempty"`
}

func (o CceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CceInfo struct{}"
	}

	return strings.Join([]string{"CceInfo", string(data)}, " ")
}
