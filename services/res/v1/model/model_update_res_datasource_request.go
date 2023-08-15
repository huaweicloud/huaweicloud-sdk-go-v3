package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateResDatasourceRequest Request Object
type UpdateResDatasourceRequest struct {

	// 数据源id
	DatasourceId string `json:"datasource_id"`

	// 工作空间id
	WorkspaceId string `json:"workspace_id"`

	Body *UpdateResDatastructRequestBody `json:"body,omitempty"`
}

func (o UpdateResDatasourceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateResDatasourceRequest struct{}"
	}

	return strings.Join([]string{"UpdateResDatasourceRequest", string(data)}, " ")
}
