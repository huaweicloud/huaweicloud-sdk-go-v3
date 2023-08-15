package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePubInfoRequestBody 创建服务号一站式请求体。
type CreatePubInfoRequestBody struct {
	PubRequestBody *PubInfoRequestBody `json:"pub_request_body"`

	PortalRequestBody *PortalInfoRequestBody `json:"portal_request_body"`

	MenuRequestBody *MenuInfoRequestBody `json:"menu_request_body"`
}

func (o CreatePubInfoRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePubInfoRequestBody struct{}"
	}

	return strings.Join([]string{"CreatePubInfoRequestBody", string(data)}, " ")
}
