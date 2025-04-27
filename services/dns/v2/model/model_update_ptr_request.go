package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePtrRequest Request Object
type UpdatePtrRequest struct {

	// 反向解析ID。
	PtrId string `json:"ptr_id"`

	Body *UpdatePtrRequestBody `json:"body,omitempty"`
}

func (o UpdatePtrRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePtrRequest struct{}"
	}

	return strings.Join([]string{"UpdatePtrRequest", string(data)}, " ")
}
