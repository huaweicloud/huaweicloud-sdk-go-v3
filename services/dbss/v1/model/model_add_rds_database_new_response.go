package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddRdsDatabaseNewResponse Response Object
type AddRdsDatabaseNewResponse struct {

	// 结果列表
	RetList        *[]RdsDbResponseRetList `json:"ret_list,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o AddRdsDatabaseNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddRdsDatabaseNewResponse struct{}"
	}

	return strings.Join([]string{"AddRdsDatabaseNewResponse", string(data)}, " ")
}
