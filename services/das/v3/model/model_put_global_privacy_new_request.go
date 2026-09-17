package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PutGlobalPrivacyNewRequest Request Object
type PutGlobalPrivacyNewRequest struct {

	// 同意状态（1：同意）
	AgreeStatus int32 `json:"agree_status"`
}

func (o PutGlobalPrivacyNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PutGlobalPrivacyNewRequest struct{}"
	}

	return strings.Join([]string{"PutGlobalPrivacyNewRequest", string(data)}, " ")
}
