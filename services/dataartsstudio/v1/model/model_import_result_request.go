package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportResultRequest Request Object
type ImportResultRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// 需要查询的某次导入的处理结果
	Uuid string `json:"uuid"`
}

func (o ImportResultRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportResultRequest struct{}"
	}

	return strings.Join([]string{"ImportResultRequest", string(data)}, " ")
}
