package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LockPortRequestBody 通道号绑定服务号请求体。
type LockPortRequestBody struct {

	// 服务号ID。
	PubId string `json:"pub_id"`

	// 通道号。 - port_type=5时，长度必须为5 - port_type=1或3，长度在21位内
	Port string `json:"port"`

	// 绑定的地区，不允许传入重叠地区。地区取值见《地区名称列表》。
	Province string `json:"province"`

	// 绑定签名，必须是该端口号签名的子集。单个签名长度为2-18。
	Sign []string `json:"sign"`

	// 关联通道号ID，取通道号列表返回的ID。
	ExtPortId string `json:"ext_port_id"`
}

func (o LockPortRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LockPortRequestBody struct{}"
	}

	return strings.Join([]string{"LockPortRequestBody", string(data)}, " ")
}
