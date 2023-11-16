package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchIdBody 批量id
type BatchIdBody struct {

	// 批量id
	Ids *[]string `json:"ids,omitempty"`
}

func (o BatchIdBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchIdBody struct{}"
	}

	return strings.Join([]string{"BatchIdBody", string(data)}, " ")
}
