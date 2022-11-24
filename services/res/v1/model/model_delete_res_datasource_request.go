package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteResDatasourceRequest struct {

	// 工作空间id。
	WorkspaceId string `json:"workspace_id"`

	// 数据源id。
	DatasourceId string `json:"datasource_id"`
}

func (o DeleteResDatasourceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteResDatasourceRequest struct{}"
	}

	return strings.Join([]string{"DeleteResDatasourceRequest", string(data)}, " ")
}
