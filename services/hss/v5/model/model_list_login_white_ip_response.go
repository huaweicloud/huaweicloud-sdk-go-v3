package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLoginWhiteIpResponse Response Object
type ListLoginWhiteIpResponse struct {

	// SSH登录IP白名单总数
	TotalNum *int32 `json:"total_num,omitempty"`

	// SSH登录IP白名单信息列表
	DataList       *[]LoginWhiteIpResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o ListLoginWhiteIpResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLoginWhiteIpResponse struct{}"
	}

	return strings.Join([]string{"ListLoginWhiteIpResponse", string(data)}, " ")
}
