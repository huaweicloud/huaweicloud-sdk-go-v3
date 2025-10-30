package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLocalImageHostsResponse Response Object
type ListLocalImageHostsResponse struct {

	// 总数
	TotalNum *int32 `json:"total_num,omitempty"`

	// 查询本地镜像的主机列表
	DataList       *[]LocalImageHostInfo `json:"data_list,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListLocalImageHostsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLocalImageHostsResponse struct{}"
	}

	return strings.Join([]string{"ListLocalImageHostsResponse", string(data)}, " ")
}
