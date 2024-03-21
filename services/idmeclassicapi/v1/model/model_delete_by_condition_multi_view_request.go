package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteByConditionMultiViewRequest Request Object
type DeleteByConditionMultiViewRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	Body *RdmParamVoDeleteByConditionVo `json:"body,omitempty"`
}

func (o DeleteByConditionMultiViewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteByConditionMultiViewRequest struct{}"
	}

	return strings.Join([]string{"DeleteByConditionMultiViewRequest", string(data)}, " ")
}
