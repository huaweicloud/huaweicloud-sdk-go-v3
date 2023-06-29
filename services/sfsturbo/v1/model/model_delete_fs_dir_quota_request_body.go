package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteFsDirQuotaRequestBody 需要删除quota的目录
type DeleteFsDirQuotaRequestBody struct {

	// 合法的已存在的目录的全路径
	Path string `json:"path"`
}

func (o DeleteFsDirQuotaRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteFsDirQuotaRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteFsDirQuotaRequestBody", string(data)}, " ")
}
