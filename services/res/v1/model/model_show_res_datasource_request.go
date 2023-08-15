package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowResDatasourceRequest Request Object
type ShowResDatasourceRequest struct {

	// 工作空间id。
	WorkspaceId string `json:"workspace_id"`

	// 数据源id。
	DatasourceId string `json:"datasource_id"`
}

func (o ShowResDatasourceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResDatasourceRequest struct{}"
	}

	return strings.Join([]string{"ShowResDatasourceRequest", string(data)}, " ")
}
