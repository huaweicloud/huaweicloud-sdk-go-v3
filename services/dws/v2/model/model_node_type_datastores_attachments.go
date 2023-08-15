package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NodeTypeDatastoresAttachments 内核版本附加信息。
type NodeTypeDatastoresAttachments struct {

	// 内核版本支持的最小CN。
	MinCn string `json:"min_cn"`

	// 内核版本支持的最大CN。
	MaxCn string `json:"max_cn"`
}

func (o NodeTypeDatastoresAttachments) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeTypeDatastoresAttachments struct{}"
	}

	return strings.Join([]string{"NodeTypeDatastoresAttachments", string(data)}, " ")
}
