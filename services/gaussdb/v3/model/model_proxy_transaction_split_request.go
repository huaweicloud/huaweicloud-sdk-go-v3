package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ProxyTransactionSplitRequest struct {

	// 开启/关闭事务拆分，取值范围是[ON/OFF]
	TransactionSplit string `json:"transaction_split"`

	// 实例的数据库代理列表，仅支持单proxy使用。
	ProxyIdList []string `json:"proxy_id_list"`
}

func (o ProxyTransactionSplitRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProxyTransactionSplitRequest struct{}"
	}

	return strings.Join([]string{"ProxyTransactionSplitRequest", string(data)}, " ")
}
