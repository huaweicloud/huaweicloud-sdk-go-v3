package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RdsDbResponseRetList struct {

	// rds数据库id
	Id string `json:"id"`

	// 状态 - SUCCESS - FAILED
	RetStatus string `json:"ret_status"`

	// 描述
	RetMessage string `json:"ret_message"`
}

func (o RdsDbResponseRetList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RdsDbResponseRetList struct{}"
	}

	return strings.Join([]string{"RdsDbResponseRetList", string(data)}, " ")
}
