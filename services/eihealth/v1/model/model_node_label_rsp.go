package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NodeLabelRsp 标签对象
type NodeLabelRsp struct {

	// 标签名称
	Name *string `json:"name,omitempty"`
}

func (o NodeLabelRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeLabelRsp struct{}"
	}

	return strings.Join([]string{"NodeLabelRsp", string(data)}, " ")
}
