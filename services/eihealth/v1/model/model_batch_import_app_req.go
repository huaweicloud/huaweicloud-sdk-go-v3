package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 导入应用请求体
type BatchImportAppReq struct {

	// 源项目id
	SourceProjectId string `json:"source_project_id"`

	// 源应用列表
	ImportApps []AppSrcReq `json:"import_apps"`
}

func (o BatchImportAppReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchImportAppReq struct{}"
	}

	return strings.Join([]string{"BatchImportAppReq", string(data)}, " ")
}
