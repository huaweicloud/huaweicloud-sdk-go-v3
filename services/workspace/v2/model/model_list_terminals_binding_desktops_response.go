package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTerminalsBindingDesktopsResponse Response Object
type ListTerminalsBindingDesktopsResponse struct {

	// MAC绑定VM信息列表
	BindList *[]TerminalsBindingDesktopsInfo `json:"bind_list,omitempty"`

	// 返回结果总条数
	TotalNum       *int32 `json:"total_num,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListTerminalsBindingDesktopsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTerminalsBindingDesktopsResponse struct{}"
	}

	return strings.Join([]string{"ListTerminalsBindingDesktopsResponse", string(data)}, " ")
}
