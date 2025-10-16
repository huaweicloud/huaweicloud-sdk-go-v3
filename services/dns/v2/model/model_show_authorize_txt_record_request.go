package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAuthorizeTxtRecordRequest Request Object
type ShowAuthorizeTxtRecordRequest struct {

	// 待创建的子域名。
	ZoneName string `json:"zone_name"`
}

func (o ShowAuthorizeTxtRecordRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAuthorizeTxtRecordRequest struct{}"
	}

	return strings.Join([]string{"ShowAuthorizeTxtRecordRequest", string(data)}, " ")
}
