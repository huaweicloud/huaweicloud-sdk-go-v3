package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTwoFactorLoginHostResponse Response Object
type ListTwoFactorLoginHostResponse struct {

	// 总数
	TotalNum *int32 `json:"total_num,omitempty"`

	// 主机列表
	DataList       *[]TwoFactorLoginHostResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                               `json:"-"`
}

func (o ListTwoFactorLoginHostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTwoFactorLoginHostResponse struct{}"
	}

	return strings.Join([]string{"ListTwoFactorLoginHostResponse", string(data)}, " ")
}
