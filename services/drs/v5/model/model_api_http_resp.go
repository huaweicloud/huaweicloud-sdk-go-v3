package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ApiHttpResp 空响应体。
type ApiHttpResp struct {
}

func (o ApiHttpResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiHttpResp struct{}"
	}

	return strings.Join([]string{"ApiHttpResp", string(data)}, " ")
}
