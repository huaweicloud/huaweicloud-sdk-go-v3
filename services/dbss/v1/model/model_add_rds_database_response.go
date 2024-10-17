package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddRdsDatabaseResponse Response Object
type AddRdsDatabaseResponse struct {

	// 结果列表
	RetList        *[]RdsDbResponseRetList `json:"ret_list,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o AddRdsDatabaseResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddRdsDatabaseResponse struct{}"
	}

	return strings.Join([]string{"AddRdsDatabaseResponse", string(data)}, " ")
}
