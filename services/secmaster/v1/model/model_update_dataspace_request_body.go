package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateDataspaceRequestBody struct {

	// 描述
	Description string `json:"description"`
}

func (o UpdateDataspaceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDataspaceRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateDataspaceRequestBody", string(data)}, " ")
}
