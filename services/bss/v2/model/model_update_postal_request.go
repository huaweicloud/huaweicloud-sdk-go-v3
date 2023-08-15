package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePostalRequest Request Object
type UpdatePostalRequest struct {

	// 语言。中文：zh_CN 英文：en_US，缺省为zh_CN
	XLanguage *string `json:"X-Language,omitempty"`

	Body *UpdatePostalReq `json:"body,omitempty"`
}

func (o UpdatePostalRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePostalRequest struct{}"
	}

	return strings.Join([]string{"UpdatePostalRequest", string(data)}, " ")
}
