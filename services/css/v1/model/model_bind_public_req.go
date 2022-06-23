package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BindPublicReq struct {
	Eip *BindPublicReqEip `json:"eip"`
}

func (o BindPublicReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BindPublicReq struct{}"
	}

	return strings.Join([]string{"BindPublicReq", string(data)}, " ")
}
