package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 内核版本信息。
type NodeTypeDatastores struct {

	// 内核版本号。
	Version string `json:"version"`

	Attachments *NodeTypeDatastoresAttachments `json:"attachments"`
}

func (o NodeTypeDatastores) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeTypeDatastores struct{}"
	}

	return strings.Join([]string{"NodeTypeDatastores", string(data)}, " ")
}
