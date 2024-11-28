package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRedisPitrInfoResponse Response Object
type ShowRedisPitrInfoResponse struct {

	// 查询Redis实例指定时间点恢复所占用的存储空间。 单位：GB
	Storage        *string `json:"storage,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowRedisPitrInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRedisPitrInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowRedisPitrInfoResponse", string(data)}, " ")
}
