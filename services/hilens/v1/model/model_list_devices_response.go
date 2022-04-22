package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListDevicesResponse struct {

	// 满足条件的设备总数
	Total *int32 `json:"total,omitempty"`

	Data           *[]GetDevicesListArrayObject `json:"data,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o ListDevicesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDevicesResponse struct{}"
	}

	return strings.Join([]string{"ListDevicesResponse", string(data)}, " ")
}
