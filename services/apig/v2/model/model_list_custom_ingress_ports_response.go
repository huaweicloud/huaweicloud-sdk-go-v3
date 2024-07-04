package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCustomIngressPortsResponse Response Object
type ListCustomIngressPortsResponse struct {

	// 本次返回的列表长度
	Size int32 `json:"size"`

	// 满足条件的记录数
	Total int64 `json:"total"`

	// 实例自定义入方向端口列表。
	IngressPortInfos *[]IngressPortInfo `json:"ingress_port_infos,omitempty"`
	HttpStatusCode   int                `json:"-"`
}

func (o ListCustomIngressPortsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCustomIngressPortsResponse struct{}"
	}

	return strings.Join([]string{"ListCustomIngressPortsResponse", string(data)}, " ")
}
