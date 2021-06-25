package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListTopStatisticsRequest struct {
	// 加速域名，格式：www.test1.com。ALL表示查询名下全部域名。（TopN视频信息要么查询单个域名要么查询所有域名）

	Domain string `json:"domain"`
	// 查询日期，格式为yyyymmdd。 1）  date必须为昨天或之前的日期 2）  只能查最近一个月内的数据

	Date string `json:"date"`
}

func (o ListTopStatisticsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListTopStatisticsRequest struct{}"
	}

	return strings.Join([]string{"ListTopStatisticsRequest", string(data)}, " ")
}
