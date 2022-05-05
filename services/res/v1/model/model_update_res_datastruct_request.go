package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateResDatastructRequest struct {

	// 数据源id。
	DatasourceId string `json:"datasource_id"`

	// 工作空间id。
	WorkspaceId string `json:"workspace_id"`

	Body *UpdateResDatastructRequestBodyBody `json:"body,omitempty"`
}

func (o UpdateResDatastructRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateResDatastructRequest struct{}"
	}

	return strings.Join([]string{"UpdateResDatastructRequest", string(data)}, " ")
}
