package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CustomizeSchemaVersionCreateReq struct {

	// 事件模型内容定义
	Definition *string `json:"definition,omitempty"`
}

func (o CustomizeSchemaVersionCreateReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomizeSchemaVersionCreateReq struct{}"
	}

	return strings.Join([]string{"CustomizeSchemaVersionCreateReq", string(data)}, " ")
}
