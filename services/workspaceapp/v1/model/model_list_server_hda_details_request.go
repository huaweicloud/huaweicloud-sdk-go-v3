package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListServerHdaDetailsRequest Request Object
type ListServerHdaDetailsRequest struct {

	// 查询的偏移量，默认值0。
	Offset *int32 `json:"offset,omitempty"`

	// 查询的数量，值区间[1-100]，默认值10。
	Limit *int32 `json:"limit,omitempty"`

	// 服务器组id。
	ServerGroupId *string `json:"server_group_id,omitempty"`

	// 服务器名称。
	ServerName *string `json:"server_name,omitempty"`
}

func (o ListServerHdaDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServerHdaDetailsRequest struct{}"
	}

	return strings.Join([]string{"ListServerHdaDetailsRequest", string(data)}, " ")
}
