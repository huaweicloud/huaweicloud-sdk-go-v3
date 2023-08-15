package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteFunctionRequest Request Object
type DeleteFunctionRequest struct {

	// 实例Id
	InstanceId string `json:"instance_id"`

	// catalog名字
	CatalogName string `json:"catalog_name"`

	// 数据库名字
	DatabaseName string `json:"database_name"`

	// 函数名字
	FunctionName string `json:"function_name"`
}

func (o DeleteFunctionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteFunctionRequest struct{}"
	}

	return strings.Join([]string{"DeleteFunctionRequest", string(data)}, " ")
}
