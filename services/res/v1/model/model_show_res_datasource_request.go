package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowResDatasourceRequest struct {

	// 工作空间id。
	WorkspaceId string `json:"workspace_id" xml:"workspace_id"`

	// 数据源id。
	DatasourceId string `json:"datasource_id" xml:"datasource_id"`
}

func (o ShowResDatasourceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResDatasourceRequest struct{}"
	}

	return strings.Join([]string{"ShowResDatasourceRequest", string(data)}, " ")
}
