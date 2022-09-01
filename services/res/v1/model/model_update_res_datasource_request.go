package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateResDatasourceRequest struct {

	// 数据源id
	DatasourceId string `json:"datasource_id" xml:"datasource_id"`

	// 工作空间id
	WorkspaceId string `json:"workspace_id" xml:"workspace_id"`

	Body *UpdateResDatastructRequestBody `json:"body,omitempty" xml:"body"`
}

func (o UpdateResDatasourceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateResDatasourceRequest struct{}"
	}

	return strings.Join([]string{"UpdateResDatasourceRequest", string(data)}, " ")
}
