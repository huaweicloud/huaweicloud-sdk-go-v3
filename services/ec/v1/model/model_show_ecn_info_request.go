package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowEcnInfoRequest Request Object
type ShowEcnInfoRequest struct {

	// 企业连接网络ID
	EcnId string `json:"ecn_id"`
}

func (o ShowEcnInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEcnInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowEcnInfoRequest", string(data)}, " ")
}
