package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateGaussMySqlInstanceResponse struct {
	Instance *MysqlInstanceResponse `json:"instance,omitempty"`
	// 实例创建的任务id。  仅创建按需实例时会返回该参数。

	JobId *string `json:"job_id,omitempty"`
	// 订单号，创建包年包月时返回该参数。

	OrderId        *string `json:"order_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateGaussMySqlInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGaussMySqlInstanceResponse struct{}"
	}

	return strings.Join([]string{"CreateGaussMySqlInstanceResponse", string(data)}, " ")
}
