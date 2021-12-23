package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 绑定弹性公网IP的请求体
type AssociatePublicipsRequestBody struct {
	Publicip *AssociatePublicipsOption `json:"publicip"`
}

func (o AssociatePublicipsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociatePublicipsRequestBody struct{}"
	}

	return strings.Join([]string{"AssociatePublicipsRequestBody", string(data)}, " ")
}
