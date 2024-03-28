package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCentralNetworkErRouteTableAttachmentsRequest Request Object
type ListCentralNetworkErRouteTableAttachmentsRequest struct {

	// 每页返回的个数。 取值范围：1~1000。
	Limit *int32 `json:"limit,omitempty"`

	// 翻页信息，从上次API调用返回的翻页数据中获取，可填写前一页marker或者后一页marker，填入前一页previous_marker就向前翻页，后一页next_marker就向翻页。 翻页过程中，查询条件不能修改，包括过滤条件，排序条件，limit。
	Marker *string `json:"marker,omitempty"`

	// 排序字段。
	SortKey *string `json:"sort_key,omitempty"`

	// 指定排序是升序还是降序(asc为升序，desc为降序)。
	SortDir *SortDir `json:"sort_dir,omitempty"`

	// 根据id查询，可查询多个id。
	Id *[]string `json:"id,omitempty"`

	// 根据名字查询，可查询多个名字。
	Name *[]string `json:"name,omitempty"`

	// 根据状态查询，可查询多个状态。
	State *[]CentralNetworkConnectionStateEnum `json:"state,omitempty"`

	// 中心网络的ID。
	CentralNetworkId string `json:"central_network_id"`

	// Attachment实例的ID。
	AttachmentInstanceId *[]string `json:"attachment_instance_id,omitempty"`
}

func (o ListCentralNetworkErRouteTableAttachmentsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCentralNetworkErRouteTableAttachmentsRequest struct{}"
	}

	return strings.Join([]string{"ListCentralNetworkErRouteTableAttachmentsRequest", string(data)}, " ")
}
