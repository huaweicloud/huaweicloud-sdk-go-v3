package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 请求id响应体
type RequestIdResponseBody struct {

	// 请求的唯一标识ID
	RequestId string `json:"request_id" xml:"request_id"`
}

func (o RequestIdResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RequestIdResponseBody struct{}"
	}

	return strings.Join([]string{"RequestIdResponseBody", string(data)}, " ")
}
