package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ClusterFlavorResp struct {

	// ID
	Id *string `json:"id,omitempty"`

	// 连接
	Links *[]LinkResp `json:"links,omitempty"`
}

func (o ClusterFlavorResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterFlavorResp struct{}"
	}

	return strings.Join([]string{"ClusterFlavorResp", string(data)}, " ")
}
