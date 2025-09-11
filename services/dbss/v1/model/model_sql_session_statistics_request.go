package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SqlSessionStatisticsRequest struct {

	// 数据库ID
	DbId *[]string `json:"db_id,omitempty"`

	Time *Duration `json:"time"`
}

func (o SqlSessionStatisticsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SqlSessionStatisticsRequest struct{}"
	}

	return strings.Join([]string{"SqlSessionStatisticsRequest", string(data)}, " ")
}
