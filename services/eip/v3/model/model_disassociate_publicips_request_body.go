package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 解绑弹性公网IP的请求体
type DisassociatePublicipsRequestBody struct {
	Publicip *DisassociatePublicipsOption `json:"publicip" xml:"publicip"`
}

func (o DisassociatePublicipsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociatePublicipsRequestBody struct{}"
	}

	return strings.Join([]string{"DisassociatePublicipsRequestBody", string(data)}, " ")
}
