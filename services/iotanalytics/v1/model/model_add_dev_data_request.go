package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type AddDevDataRequest struct {

	// 数据源id
	DatasourceId string `json:"datasource_id" xml:"datasource_id"`

	Body *interface{} `json:"body,omitempty" xml:"body"`
}

func (o AddDevDataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddDevDataRequest struct{}"
	}

	return strings.Join([]string{"AddDevDataRequest", string(data)}, " ")
}
