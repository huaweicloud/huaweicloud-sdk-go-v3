package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateParametersReq struct {
	Values *UpdateParametersReqValues `json:"values"`
}

func (o UpdateParametersReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateParametersReq struct{}"
	}

	return strings.Join([]string{"UpdateParametersReq", string(data)}, " ")
}
