package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 弹性Ip信息
type CreateClusterPublicIpReq struct {
	Eip *CreateClusterPublicEip `json:"eip"`

	ElbWhiteListReq *CreateClusterElbWhiteList `json:"elbWhiteListReq"`
	// 是否自动绑定弹性公网IP。auto_assign为自动分配，bind_existing为绑定已有IP，需要填写eipId字段。

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
