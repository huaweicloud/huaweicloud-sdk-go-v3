package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateFunctionRequest struct {

	// 实例Id
	InstanceId string `json:"instance_id"`

	// catalog名字
	CatalogName string `json:"catalog_name"`

	// 数据库名字
	DatabaseName string `json:"database_name"`

	Body *FunctionInput `json:"body,omitempty"`
}

func (o CreateFunctionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFunctionRequest struct{}"
	}

	return strings.Join([]string{"CreateFunctionRequest", string(data)}, " ")
}
