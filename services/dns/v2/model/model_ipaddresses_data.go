package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IpaddressesData struct {

	// 资源状态。
	Status *string `json:"status,omitempty"`

	// 终端节点的ID，UUID形式的一个资源标识。
	Id *string `json:"id,omitempty"`

	// IP地址信息。
	Ip *string `json:"ip,omitempty"`

	// 创建时间。 格式：yyyy-MM-dd'T'HH:mm:ss.SSS。
	CreateTime *string `json:"create_time,omitempty"`

	// 更新时间。 格式：yyyy-MM-dd'T'HH:mm:ss.SSS。
	UpdateTime *string `json:"update_time,omitempty"`

	// 子网ID。
	SubnetId *string `json:"subnet_id,omitempty"`

	// 错误信息。
	ErrorInfo *string `json:"error_info,omitempty"`
}

func (o IpaddressesData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IpaddressesData struct{}"
	}

	return strings.Join([]string{"IpaddressesData", string(data)}, " ")
}
