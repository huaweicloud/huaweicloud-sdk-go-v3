package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDataSourceResponse Response Object
type CreateDataSourceResponse struct {

	// 数据源id。
	Id *string `json:"id,omitempty"`

	// 创建数据源job_id。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateDataSourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDataSourceResponse struct{}"
	}

	return strings.Join([]string{"CreateDataSourceResponse", string(data)}, " ")
}
