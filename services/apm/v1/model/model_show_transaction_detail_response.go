package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTransactionDetailResponse Response Object
type ShowTransactionDetailResponse struct {

	// 组件节点列表。
	TxNodeList *[]TxNode `json:"tx_node_list,omitempty"`

	// 组件之间调用指向线列表。
	TxLineList     *[]TxLine `json:"tx_line_list,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowTransactionDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTransactionDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowTransactionDetailResponse", string(data)}, " ")
}
