package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListConnectionResponse Response Object
type ListConnectionResponse struct {

	// 总记录数目
	TotalRecord *int32 `json:"total_record,omitempty"`

	// 连接信息列表
	DasConnInfoList *[]DasConnInfo `json:"das_conn_info_list,omitempty"`
	HttpStatusCode  int            `json:"-"`
}

func (o ListConnectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConnectionResponse struct{}"
	}

	return strings.Join([]string{"ListConnectionResponse", string(data)}, " ")
}
