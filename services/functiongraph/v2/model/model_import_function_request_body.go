package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 导入函数请求结构体
type ImportFunctionRequestBody struct {
	// 函数名

	FuncName string `json:"func_name"`
	// 文件名

	FileName string `json:"file_name"`
	// 文件类型

	FileType string `json:"file_type"`
	// 函数代码。代码必须要进行base64编码

	FileCode string `json:"file_code"`
}

func (o ImportFunctionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportFunctionRequestBody struct{}"
	}

	return strings.Join([]string{"ImportFunctionRequestBody", string(data)}, " ")
}
