package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowClusterListRequest Request Object
type ShowClusterListRequest struct {

	// 获取特定category的集群。
	Category *string `json:"category,omitempty"`

	// 是否获取集群的资源信息。不填或者填写为true为获取集群资源汇总信息，置为false为不获取集群状态信息；缺省值为true
	Enablestatus *bool `json:"enablestatus,omitempty"`

	// 容器舰队ID。不填会返回用户所有集群，填了之后会返回属于该容器舰队的集群。
	Clustergroupid *string `json:"clustergroupid,omitempty"`

	// 分页获取列表时，页的大小，默认为-1
	Limit *int32 `json:"limit,omitempty"`

	// 分页获取列表时，起始偏移量，默认为0
	Offset *int32 `json:"offset,omitempty"`

	// 分页获取列表时，排序参数，支持 create_at 和 update_at
	OrderBy *string `json:"order_by,omitempty"`

	// 分页获取列表时，排序方向，支持 desc 和 asc
	Order *string `json:"order,omitempty"`

	// 获取集群列表时，根据集群类型筛选，不传参时默认为 all ，支持 all ，grouped，discrete 三种类型。 - grouped：在舰队中纳管的集群 - discrete：未加入舰队的集群 - all：所有集群
	Managetype *string `json:"managetype,omitempty"`

	// 集群ID。多个ID以英文逗号分隔。
	Clusterids *string `json:"clusterids,omitempty"`
}

func (o ShowClusterListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowClusterListRequest struct{}"
	}

	return strings.Join([]string{"ShowClusterListRequest", string(data)}, " ")
}
