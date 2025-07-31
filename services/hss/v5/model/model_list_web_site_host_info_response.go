package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWebSiteHostInfoResponse Response Object
type ListWebSiteHostInfoResponse struct {

	// 总数
	TotalNum *int32 `json:"total_num,omitempty"`

	// 服务器列表
	DataList       *[]WebSiteHostInfo `json:"data_list,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListWebSiteHostInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWebSiteHostInfoResponse struct{}"
	}

	return strings.Join([]string{"ListWebSiteHostInfoResponse", string(data)}, " ")
}
