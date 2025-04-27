package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPtrRequest Request Object
type ShowPtrRequest struct {

	// 反向解析ID。
	PtrId string `json:"ptr_id"`
}

func (o ShowPtrRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPtrRequest struct{}"
	}

	return strings.Join([]string{"ShowPtrRequest", string(data)}, " ")
}
