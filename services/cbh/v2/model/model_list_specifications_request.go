package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSpecificationsRequest Request Object
type ListSpecificationsRequest struct {

	// 查询云堡垒机规格当前动作。 - create：查询可创建云堡垒机规格信息 - update：查询可变更云堡垒机规格信息
	Action string `json:"action"`

	// 云堡垒机规格信息，当action为update时此字段必填。
	SpecCode *string `json:"spec_code,omitempty"`
}

func (o ListSpecificationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSpecificationsRequest struct{}"
	}

	return strings.Join([]string{"ListSpecificationsRequest", string(data)}, " ")
}
