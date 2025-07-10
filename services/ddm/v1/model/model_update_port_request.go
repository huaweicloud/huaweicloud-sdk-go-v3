package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdatePortRequest struct {

	// 目标端口号
	Port int32 `json:"port"`
}

func (o UpdatePortRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePortRequest struct{}"
	}

	return strings.Join([]string{"UpdatePortRequest", string(data)}, " ")
}
