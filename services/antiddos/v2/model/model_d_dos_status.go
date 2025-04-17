package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DDosStatus ddos状态响应体
type DDosStatus struct {

	// EIP的ID
	FloatingIpId string `json:"floating_ip_id"`

	// 浮动IP地址
	FloatingIpAddress string `json:"floating_ip_address"`

	// EIP所属类型，可选范围： - Anti-DDoS：eip属于Anti-DDoS流量清洗 - CNAD：eip属于DDoS原生高级防护
	ProductType string `json:"product_type"`

	// 防护状态，可选范围： - normal：表示正常 - configging：表示设置中 - notConfig：表示未设置 - packetcleaning：表示清洗 - packetdropping：表示黑洞
	Status string `json:"status"`

	// 清洗阈值
	CleanThreshold int64 `json:"clean_threshold"`

	// 黑洞阈值
	BlockThreshold string `json:"block_threshold"`
}

func (o DDosStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DDosStatus struct{}"
	}

	return strings.Join([]string{"DDosStatus", string(data)}, " ")
}
