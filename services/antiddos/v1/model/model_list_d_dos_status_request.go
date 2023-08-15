package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDDosStatusRequest Request Object
type ListDDosStatusRequest struct {

	// 可选范围： - normal：表示正常 - configging：表示设置中 - notConfig：表示未设置 - packetcleaning：表示清洗 - packetdropping：表示黑洞  不带此参数默认所有列表，以neutron查询到的顺序为准。
	Status *string `json:"status,omitempty"`

	// 返回结果个数限制，取值范围：1~100
	Limit *string `json:"limit,omitempty"`

	// 偏移量，取值范围：0~2147483647
	Offset *string `json:"offset,omitempty"`

	// IP地址，支持IPv4格式和IPv6格式输入，支持部分查询。例如“？ip=192.168”，会返回192.168.111.1和10.192.168.8所对应的EIP防护状态。
	Ip *string `json:"ip,omitempty"`
}

func (o ListDDosStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDDosStatusRequest struct{}"
	}

	return strings.Join([]string{"ListDDosStatusRequest", string(data)}, " ")
}
