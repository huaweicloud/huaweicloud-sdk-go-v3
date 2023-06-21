package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 请求成功返回的数据。
type UpdatePubInfoResponseModelData struct {

	// 服务号更新记录ID。
	PubUpdateLogId *string `json:"pub_update_log_id,omitempty"`
}

func (o UpdatePubInfoResponseModelData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePubInfoResponseModelData struct{}"
	}

	return strings.Join([]string{"UpdatePubInfoResponseModelData", string(data)}, " ")
}
