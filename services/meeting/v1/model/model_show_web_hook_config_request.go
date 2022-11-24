package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowWebHookConfigRequest struct {

	// 企业ID。按企业注册回调时需要填写。
	CorpId *string `json:"corpId,omitempty"`

	// SP ID。多租户场景下，按SP注册回调时需要填写。
	SpId *string `json:"spId,omitempty"`
}

func (o ShowWebHookConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWebHookConfigRequest struct{}"
	}

	return strings.Join([]string{"ShowWebHookConfigRequest", string(data)}, " ")
}
