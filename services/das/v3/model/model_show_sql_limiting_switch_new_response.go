package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSqlLimitingSwitchNewResponse Response Object
type ShowSqlLimitingSwitchNewResponse struct {

	// 开关状态
	SwitchOn       *string `json:"switch_on,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowSqlLimitingSwitchNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSqlLimitingSwitchNewResponse struct{}"
	}

	return strings.Join([]string{"ShowSqlLimitingSwitchNewResponse", string(data)}, " ")
}
