package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 拓扑树节点信息
type TreeNode struct {

	// 拓扑树节点id
	Id *string `json:"id,omitempty" xml:"id"`

	// 拓扑树节点的父节点
	Parent *string `json:"parent,omitempty" xml:"parent"`

	// 拓扑树节点的实际id
	RealId *int64 `json:"real_id,omitempty" xml:"real_id"`

	// 拓扑树节点名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 拓扑树节点展示名称
	DisplayName *string `json:"display_name,omitempty" xml:"display_name"`

	// 组件名称
	AppName *string `json:"app_name,omitempty" xml:"app_name"`

	// 组件id
	AppId *int64 `json:"app_id,omitempty" xml:"app_id"`

	// 是否是管理节点
	IsAdmin *bool `json:"is_admin,omitempty" xml:"is_admin"`

	// 是否是根节点
	IsRoot *bool `json:"is_root,omitempty" xml:"is_root"`

	// 应用id
	BusinessId *int64 `json:"business_id,omitempty" xml:"business_id"`

	// 节点类型
	NodeType *TreeNodeNodeType `json:"node_type,omitempty" xml:"node_type"`

	// 区域
	Region *string `json:"region,omitempty" xml:"region"`

	// 是否是默认的节点
	IsDefault *bool `json:"is_default,omitempty" xml:"is_default"`
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
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
