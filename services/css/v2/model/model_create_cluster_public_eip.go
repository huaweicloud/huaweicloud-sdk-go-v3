package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 弹性IP白名单信息
type CreateClusterPublicEip struct {
	BandWidth *CreateClusterPublicEipSize `json:"bandWidth"`
}

func (o CreateClusterPublicEip) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateClusterPublicEip struct{}"
	}

	return strings.Join([]string{"CreateClusterPublicEip", string(data)}, " ")
}
