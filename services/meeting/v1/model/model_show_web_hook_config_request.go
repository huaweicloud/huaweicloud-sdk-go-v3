package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowWebHookConfigRequest struct {
	// 企业ID，与spId二者必填一个

	CorpId *string `json:"corpId,omitempty"`
	// sp管理员ID，与corpId二者必填一个

	SpId *string `json:"spId,omitempty"`
}

func (o ShowWebHookConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWebHookConfigRequest struct{}"
	}

	return strings.Join([]string{"ShowWebHookConfigRequest", string(data)}, " ")
}
