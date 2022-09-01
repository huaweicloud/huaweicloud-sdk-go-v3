package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Kibana公网访问信息
type CreateClusterPublicKibanaReq struct {

	// 带宽大小。
	EipSize int32 `json:"eipSize" xml:"eipSize"`

	ElbWhiteList *CreateClusterPublicKibanaElbWhiteList `json:"elbWhiteList" xml:"elbWhiteList"`
}

func (o CreateClusterPublicKibanaReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateClusterPublicKibanaReq struct{}"
	}

	return strings.Join([]string{"CreateClusterPublicKibanaReq", string(data)}, " ")
}
