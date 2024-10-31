package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHttpIpGroupsResponse Response Object
type ShowHttpIpGroupsResponse struct {

	// 全部IP地址组的数量
	Total *int32 `json:"total,omitempty"`

	// 详细的IP地址组信息
	Items          *[]ShowHttpIpGroupResponseBody `json:"items,omitempty"`
	HttpStatusCode int                            `json:"-"`
}

func (o ShowHttpIpGroupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHttpIpGroupsResponse struct{}"
	}

	return strings.Join([]string{"ShowHttpIpGroupsResponse", string(data)}, " ")
}
