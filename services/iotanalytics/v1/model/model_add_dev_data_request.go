package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type AddDevDataRequest struct {
	// 数据源id

	DatasourceId string `json:"datasource_id"`

	Body *interface{} `json:"body,omitempty"`
}

func (o AddDevDataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddDevDataRequest struct{}"
	}

	return strings.Join([]string{"AddDevDataRequest", string(data)}, " ")
}
