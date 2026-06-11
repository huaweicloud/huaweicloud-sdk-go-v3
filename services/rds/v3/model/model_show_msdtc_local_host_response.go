package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMsdtcLocalHostResponse Response Object
type ShowMsdtcLocalHostResponse struct {

	// 查询状态 processing:查询中 success：查询成功 fail:查询失败
	Status *string `json:"status,omitempty"`

	// host信息列表
	Hosts          *[]MsdtcHostResult `json:"hosts,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ShowMsdtcLocalHostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMsdtcLocalHostResponse struct{}"
	}

	return strings.Join([]string{"ShowMsdtcLocalHostResponse", string(data)}, " ")
}
