package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDatabaseDataRequest Request Object
type DeleteDatabaseDataRequest struct {

	// 平台项目ID，您可以在平台单击所需的项目名称，进入项目设置页面查看。
	EihealthProjectId string `json:"eihealth_project_id"`

	// 数据库实例id
	DatabaseId string `json:"database_id"`

	// 数据的行号，即_row_num值
	RowNum int64 `json:"row_num"`
}

func (o DeleteDatabaseDataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDatabaseDataRequest struct{}"
	}

	return strings.Join([]string{"DeleteDatabaseDataRequest", string(data)}, " ")
}
