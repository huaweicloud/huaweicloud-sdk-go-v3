package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateDatabaseDataRequest struct {

	// 医疗智能体平台项目ID，您可以在EIHealth平台单击所需的项目名称，进入项目设置页面查看。
	EihealthProjectId string `json:"eihealth_project_id"`

	// 数据库实例id
	DatabaseId string `json:"database_id"`

	// 数据行号，即_row_num值
	RowNum int64 `json:"row_num"`

	Body *RowDataReq `json:"body,omitempty"`
}

func (o UpdateDatabaseDataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDatabaseDataRequest struct{}"
	}

	return strings.Join([]string{"UpdateDatabaseDataRequest", string(data)}, " ")
}
