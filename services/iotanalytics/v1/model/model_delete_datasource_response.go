package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteDatasourceResponse struct {

	// 删除数据源的id
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteDatasourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDatasourceResponse struct{}"
	}

	return strings.Join([]string{"DeleteDatasourceResponse", string(data)}, " ")
}
