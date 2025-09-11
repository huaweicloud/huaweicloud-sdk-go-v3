package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ChangeNameRequest struct {

	// 新名称
	NewName string `json:"new_name"`
}

func (o ChangeNameRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeNameRequest struct{}"
	}

	return strings.Join([]string{"ChangeNameRequest", string(data)}, " ")
}
