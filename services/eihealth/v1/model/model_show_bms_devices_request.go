package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBmsDevicesRequest Request Object
type ShowBmsDevicesRequest struct {

	// 计算资源id
	Id string `json:"id"`
}

func (o ShowBmsDevicesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBmsDevicesRequest struct{}"
	}

	return strings.Join([]string{"ShowBmsDevicesRequest", string(data)}, " ")
}
