package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Agency struct {

	// 委托名称。
	AgencyName string `json:"agency_name"`

	// 委托ID。
	AgencyId *string `json:"agency_id,omitempty"`

	// 异常信息。当委托创建失败时，返回的异常信息。
	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o Agency) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Agency struct{}"
	}

	return strings.Join([]string{"Agency", string(data)}, " ")
}
