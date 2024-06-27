package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResizeClickHouseFlavorResponse Response Object
type ResizeClickHouseFlavorResponse struct {

	// 请求结果。
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ResizeClickHouseFlavorResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeClickHouseFlavorResponse struct{}"
	}

	return strings.Join([]string{"ResizeClickHouseFlavorResponse", string(data)}, " ")
}
