package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HttpListAlarmWhiteListResponseData struct {

	// 查询告警白名单返回值数据
	List *[]EipInfo `json:"list,omitempty"`

	// 目前页数
	Pages *int32 `json:"pages,omitempty"`

	// 每页个数
	Size *int32 `json:"size,omitempty"`
}

func (o HttpListAlarmWhiteListResponseData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HttpListAlarmWhiteListResponseData struct{}"
	}

	return strings.Join([]string{"HttpListAlarmWhiteListResponseData", string(data)}, " ")
}
