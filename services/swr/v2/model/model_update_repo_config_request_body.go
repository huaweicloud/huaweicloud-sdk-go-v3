package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateRepoConfigRequestBody struct {

	// 制品仓库描述
	Description string `json:"description"`
}

func (o UpdateRepoConfigRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRepoConfigRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateRepoConfigRequestBody", string(data)}, " ")
}
