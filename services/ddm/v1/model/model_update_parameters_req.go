package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateParametersReq struct {
	// 需要修改的DDM实例参数的集合。

	Values *interface{} `json:"values"`
}

func (o UpdateParametersReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateParametersReq struct{}"
	}

	return strings.Join([]string{"UpdateParametersReq", string(data)}, " ")
}
