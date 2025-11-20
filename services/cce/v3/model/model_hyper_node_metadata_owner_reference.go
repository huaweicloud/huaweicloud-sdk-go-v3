package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HyperNodeMetadataOwnerReference 属主对象
type HyperNodeMetadataOwnerReference struct {

	// **参数解释**： 节点池名称
	NodepoolName *string `json:"nodepoolName,omitempty"`

	// **参数解释**： 节点池UID
	NodepoolID *string `json:"nodepoolID,omitempty"`
}

func (o HyperNodeMetadataOwnerReference) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HyperNodeMetadataOwnerReference struct{}"
	}

	return strings.Join([]string{"HyperNodeMetadataOwnerReference", string(data)}, " ")
}
