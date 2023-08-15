package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListServersDedicatedHostRequest Request Object
type ListServersDedicatedHostRequest struct {

	// 专属主机ID。  可以从专属主机控制台查询，或者通过调用查询专属主机列表API获取。
	DedicatedHostId string `json:"dedicated_host_id"`

	// 每个页面上显示的条目数。
	Limit *int32 `json:"limit,omitempty"`

	// 该值是上一页最后一条记录的ID。  如果“marker”取值无效，将会返回“400”错误码。
	Marker *string `json:"marker,omitempty"`
}

func (o ListServersDedicatedHostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServersDedicatedHostRequest struct{}"
	}

	return strings.Join([]string{"ListServersDedicatedHostRequest", string(data)}, " ")
}
