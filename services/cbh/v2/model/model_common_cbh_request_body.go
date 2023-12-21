package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CommonCbhRequestBody 云堡垒机实例请求对象。
type CommonCbhRequestBody struct {

	// 实例id
	ServerId string `json:"server_id"`
}

func (o CommonCbhRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommonCbhRequestBody struct{}"
	}

	return strings.Join([]string{"CommonCbhRequestBody", string(data)}, " ")
}
