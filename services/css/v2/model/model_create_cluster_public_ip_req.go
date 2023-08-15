package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateClusterPublicIpReq 公网访问信息。只有在httpsEnable设置为true时该参数配置生效。
type CreateClusterPublicIpReq struct {
	Eip *CreateClusterPublicEip `json:"eip"`

	ElbWhiteListReq *CreateClusterElbWhiteList `json:"elbWhiteListReq"`

	// 是否自动绑定弹性公网IP。当前仅支持auto_assign为自动分配参数。
	PublicBindType string `json:"publicBindType"`

	// 弹性公网IP的ID。
	EipId *string `json:"eipId,omitempty"`
}

func (o CreateClusterPublicIpReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateClusterPublicIpReq struct{}"
	}

	return strings.Join([]string{"CreateClusterPublicIpReq", string(data)}, " ")
}
