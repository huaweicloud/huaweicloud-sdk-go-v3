package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceDiskResponse Response Object
type ShowInstanceDiskResponse struct {

	// 已使用量。表示当前实例已使用的存储空间大小。单位：GB
	Used *string `json:"used,omitempty"`

	// 总量。表示当前实例最大存储空间大小。单位：GB
	Total          *string `json:"total,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowInstanceDiskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceDiskResponse struct{}"
	}

	return strings.Join([]string{"ShowInstanceDiskResponse", string(data)}, " ")
}
