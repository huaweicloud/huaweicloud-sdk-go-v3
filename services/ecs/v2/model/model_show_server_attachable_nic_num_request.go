package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowServerAttachableNicNumRequest Request Object
type ShowServerAttachableNicNumRequest struct {

	// 云服务器ID。
	ServerId string `json:"server_id"`
}

func (o ShowServerAttachableNicNumRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowServerAttachableNicNumRequest struct{}"
	}

	return strings.Join([]string{"ShowServerAttachableNicNumRequest", string(data)}, " ")
}
