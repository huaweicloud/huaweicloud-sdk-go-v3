package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePtrRequest Request Object
type DeletePtrRequest struct {

	// 反向解析ID。
	PtrId string `json:"ptr_id"`
}

func (o DeletePtrRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePtrRequest struct{}"
	}

	return strings.Join([]string{"DeletePtrRequest", string(data)}, " ")
}
