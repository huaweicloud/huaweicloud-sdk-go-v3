package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowVpcepConnectionResponse struct {
	Connections *[]Connections `json:"connections,omitempty" xml:"connections"`

	// 终端节点更新开关。 - true： 开启终端节点。 - false： 关闭终端节点。
	VpcepUpdateSwitch *bool `json:"vpcepUpdateSwitch,omitempty" xml:"vpcepUpdateSwitch"`

	// 终端节点数量。
	TotalCount     *int32 `json:"total_count,omitempty" xml:"total_count"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowVpcepConnectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVpcepConnectionResponse struct{}"
	}

	return strings.Join([]string{"ShowVpcepConnectionResponse", string(data)}, " ")
}
