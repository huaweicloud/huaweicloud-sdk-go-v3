package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 公网带宽信息。
type CreateClusterPublicEipSize struct {
	// 带宽大小。

	Size int32 `json:"size"`
}

func (o CreateClusterPublicEipSize) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateClusterPublicEipSize struct{}"
	}

	return strings.Join([]string{"CreateClusterPublicEipSize", string(data)}, " ")
}
