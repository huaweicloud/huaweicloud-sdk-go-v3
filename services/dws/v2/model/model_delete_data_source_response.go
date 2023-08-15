package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDataSourceResponse Response Object
type DeleteDataSourceResponse struct {

	// 删除数据源job_id。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteDataSourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDataSourceResponse struct{}"
	}

	return strings.Join([]string{"DeleteDataSourceResponse", string(data)}, " ")
}
