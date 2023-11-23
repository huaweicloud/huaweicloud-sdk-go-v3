package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateDataspaceRequestBody struct {

	// 数据空间名称
	DataspaceName string `json:"dataspace_name"`

	// 描述
	Description string `json:"description"`
}

func (o CreateDataspaceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDataspaceRequestBody struct{}"
	}

	return strings.Join([]string{"CreateDataspaceRequestBody", string(data)}, " ")
}
