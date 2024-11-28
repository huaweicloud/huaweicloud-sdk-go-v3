package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteMsdtcLocalHostRequestBody struct {

	// 删除msdtc中的host
	Hosts *[]MsdtcHostOption `json:"hosts,omitempty"`
}

func (o DeleteMsdtcLocalHostRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteMsdtcLocalHostRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteMsdtcLocalHostRequestBody", string(data)}, " ")
}
