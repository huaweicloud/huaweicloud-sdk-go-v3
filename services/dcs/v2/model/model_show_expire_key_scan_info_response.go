package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowExpireKeyScanInfoResponse Response Object
type ShowExpireKeyScanInfoResponse struct {

	// 记录
	Records *[]SimpleKeyScanRecord `json:"records,omitempty"`

	// 实例ID
	InstanceId *string `json:"instance_id,omitempty"`

	// 总数
	Total *int32 `json:"total,omitempty"`

	// 统计
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowExpireKeyScanInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowExpireKeyScanInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowExpireKeyScanInfoResponse", string(data)}, " ")
}
