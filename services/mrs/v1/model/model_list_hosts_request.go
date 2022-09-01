package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListHostsRequest struct {

	// 集群ID。获取方法，请参见[获取集群ID](https://support.huaweicloud.com/api-mrs/mrs_02_9001.html)。
	ClusterId string `json:"cluster_id" xml:"cluster_id"`

	// 分页查询每页返回的最大集群数量。 取值范围：[1～2147483646] 默认值为10。
	PageSize *string `json:"pageSize,omitempty" xml:"pageSize"`

	// 当前查询页码。默认值为1。
	CurrentPage *string `json:"currentPage,omitempty" xml:"currentPage"`
}

func (o ListHostsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHostsRequest struct{}"
	}

	return strings.Join([]string{"ListHostsRequest", string(data)}, " ")
}
