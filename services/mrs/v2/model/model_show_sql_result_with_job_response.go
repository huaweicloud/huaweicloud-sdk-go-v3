package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSqlResultWithJobResponse Response Object
type ShowSqlResultWithJobResponse struct {

	// SQL语句查询结果。
	SqlResults     *interface{} `json:"sql_results,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowSqlResultWithJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSqlResultWithJobResponse struct{}"
	}

	return strings.Join([]string{"ShowSqlResultWithJobResponse", string(data)}, " ")
}
