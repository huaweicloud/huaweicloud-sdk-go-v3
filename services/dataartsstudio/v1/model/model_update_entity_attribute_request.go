package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateEntityAttributeRequest Request Object
type UpdateEntityAttributeRequest struct {

	// 工作空间ID，获取方法请参见[实例ID和工作空间ID](dataartsstudio_02_0350.xml)。
	Workspace string `json:"workspace"`

	// 资产guid，获取方法请参见[数据资产guid](dataartsstudio_02_0351.xml)。
	Guid string `json:"guid"`

	// 要修改的属性名称，如description、alias、comment等。
	AttrName string `json:"attr_name"`

	// 要修改的属性值。
	AttrValue string `json:"attr_value"`
}

func (o UpdateEntityAttributeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEntityAttributeRequest struct{}"
	}

	return strings.Join([]string{"UpdateEntityAttributeRequest", string(data)}, " ")
}
