package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UnBindPublicReq struct {
	Eip *UnBindPublicReqEipReq `json:"eip,omitempty"`
}

func (o UnBindPublicReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnBindPublicReq struct{}"
	}

	return strings.Join([]string{"UnBindPublicReq", string(data)}, " ")
}
