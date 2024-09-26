package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RequestId 请求ID。
type RequestId struct {

	// 请求ID。
	RequestId string `json:"request_id"`
}

func (o RequestId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RequestId struct{}"
	}

	return strings.Join([]string{"RequestId", string(data)}, " ")
}
