package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ProxyTransactionSplitRequest struct {

	// 开启/关闭事务拆分，取值范围是[ON/OFF]
	TransactionSplit string `json:"transaction_split"`

	// 实例的proxy列表
	ProxyIdList []string `json:"proxy_id_list"`
}

func (o ProxyTransactionSplitRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProxyTransactionSplitRequest struct{}"
	}

	return strings.Join([]string{"ProxyTransactionSplitRequest", string(data)}, " ")
}
