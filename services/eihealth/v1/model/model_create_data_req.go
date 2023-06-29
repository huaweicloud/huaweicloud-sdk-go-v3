package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDataReq 创建文件夹请求体
type CreateDataReq struct {

	// 文件夹名称
	Name string `json:"name"`

	// 所在文件夹
	ParentFolder *string `json:"parent_folder,omitempty"`
}

func (o CreateDataReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDataReq struct{}"
	}

	return strings.Join([]string{"CreateDataReq", string(data)}, " ")
}
