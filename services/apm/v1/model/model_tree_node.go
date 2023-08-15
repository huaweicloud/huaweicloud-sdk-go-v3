package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// TreeNode 拓扑树节点信息。
type TreeNode struct {

	// 拓扑树节点id。
	Id *string `json:"id,omitempty"`

	// 拓扑树节点的父节点。
	Parent *string `json:"parent,omitempty"`

	// 拓扑树节点的实际id。
	RealId *int64 `json:"real_id,omitempty"`

	// 拓扑树节点名称。
	Name *string `json:"name,omitempty"`

	// 拓扑树节点展示名称。
	DisplayName *string `json:"display_name,omitempty"`

	// 组件名称。
	AppName *string `json:"app_name,omitempty"`

	// 组件id。
	AppId *int64 `json:"app_id,omitempty"`

	// 是否是管理节点。
	IsAdmin *bool `json:"is_admin,omitempty"`

	// 是否是根节点。
	IsRoot *bool `json:"is_root,omitempty"`

	// 应用id。
	BusinessId *int64 `json:"business_id,omitempty"`

	// 节点类型。
	NodeType *TreeNodeNodeType `json:"node_type,omitempty"`

	// 区域。
	Region *string `json:"region,omitempty"`

	// 是否是默认的节点。
	IsDefault *bool `json:"is_default,omitempty"`
}

func (o TreeNode) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TreeNode struct{}"
	}

	return strings.Join([]string{"TreeNode", string(data)}, " ")
}

type TreeNodeNodeType struct {
	value string
}

type TreeNodeNodeTypeEnum struct {
	BUSINESS     TreeNodeNodeType
	SUB_BUSINESS TreeNodeNodeType
	APPLICATION  TreeNodeNodeType
	ENVIRONMENT  TreeNodeNodeType
}

func GetTreeNodeNodeTypeEnum() TreeNodeNodeTypeEnum {
	return TreeNodeNodeTypeEnum{
		BUSINESS: TreeNodeNodeType{
			value: "BUSINESS",
		},
		SUB_BUSINESS: TreeNodeNodeType{
			value: "SUB_BUSINESS",
		},
		APPLICATION: TreeNodeNodeType{
			value: "APPLICATION",
		},
		ENVIRONMENT: TreeNodeNodeType{
			value: "ENVIRONMENT",
		},
	}
}

func (c TreeNodeNodeType) Value() string {
	return c.value
}

func (c TreeNodeNodeType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TreeNodeNodeType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
