package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDevicesInProductResponse Response Object
type ListDevicesInProductResponse struct {

	// 产品内设备数量
	Summary        *int32 `json:"summary,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListDevicesInProductResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDevicesInProductResponse struct{}"
	}

	return strings.Join([]string{"ListDevicesInProductResponse", string(data)}, " ")
}
