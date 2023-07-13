package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ComCce CCE组件
type ComCce struct {
	Cluster *Detail `json:"cluster,omitempty"`

	Network *Detail `json:"network,omitempty"`

	SecurityGroup *Detail `json:"security_group,omitempty"`
}

func (o ComCce) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ComCce struct{}"
	}

	return strings.Join([]string{"ComCce", string(data)}, " ")
}
