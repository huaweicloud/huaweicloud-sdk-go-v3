package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Kibana公网访问信息
type CreateClusterPublicKibanaReq struct {
	// 带宽大小。

	EipSize int32 `json:"eipSize"`

	ElbWhiteList *CreateClusterPublicKibanaElbWhiteList `json:"elbWhiteList"`
}

func (o CreateClusterPublicKibanaReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateClusterPublicKibanaReq struct{}"
	}

	return strings.Join([]string{"CreateClusterPublicKibanaReq", string(data)}, " ")
}
