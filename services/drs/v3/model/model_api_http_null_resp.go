package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ApiHttpNullResp 空响应体。
type ApiHttpNullResp struct {
}

func (o ApiHttpNullResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiHttpNullResp struct{}"
	}

	return strings.Join([]string{"ApiHttpNullResp", string(data)}, " ")
}
