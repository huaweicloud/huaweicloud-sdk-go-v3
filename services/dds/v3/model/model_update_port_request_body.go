package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdatePortRequestBody struct {

	// 新端口号。端口号有效范围为2100~9500，27017，27018，27019。
	Port int32 `json:"port"`
}

func (o UpdatePortRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePortRequestBody struct{}"
	}

	return strings.Join([]string{"UpdatePortRequestBody", string(data)}, " ")
}
