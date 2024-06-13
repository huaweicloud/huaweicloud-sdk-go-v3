package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowIfTaskNameRepeatResponse Response Object
type ShowIfTaskNameRepeatResponse struct {

	// 当前页
	PageNo *int32 `json:"pageNo,omitempty"`

	// 每页大小
	PageSize *int32 `json:"pageSize,omitempty"`

	// 总页数
	TotalPage *int32 `json:"totalPage,omitempty"`

	// 总条数
	TotalSize *int32 `json:"totalSize,omitempty"`

	// 查询到的告警模板
	PageList       *[]AlarmTemplateInfo `json:"pageList,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ShowIfTaskNameRepeatResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIfTaskNameRepeatResponse struct{}"
	}

	return strings.Join([]string{"ShowIfTaskNameRepeatResponse", string(data)}, " ")
}
