package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListConnectionsResponse Response Object
type ListConnectionsResponse struct {

	// 总记录数目
	TotalRecord *int32 `json:"total_record,omitempty"`

	// 连接信息列表
	DasConnInfoList *[]ApiListConnectionsInfoRespDasConnInfoList `json:"das_conn_info_list,omitempty"`
	HttpStatusCode  int                                          `json:"-"`
}

func (o ListConnectionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConnectionsResponse struct{}"
	}

	return strings.Join([]string{"ListConnectionsResponse", string(data)}, " ")
}
