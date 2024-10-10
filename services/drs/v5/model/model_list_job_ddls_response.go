package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListJobDdlsResponse Response Object
type ListJobDdlsResponse struct {

	// DDL告警信息列表。
	DdlList *[]DdlAlarmResp `json:"ddl_list,omitempty"`

	// 列表中的项目总数，与分页无关。
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListJobDdlsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListJobDdlsResponse struct{}"
	}

	return strings.Join([]string{"ListJobDdlsResponse", string(data)}, " ")
}
