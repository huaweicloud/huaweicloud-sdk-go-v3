package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteByConditionVo struct {
	Condition *QueryRequestVo `json:"condition"`

	// 修改人。
	Modifier *string `json:"modifier,omitempty"`
}

func (o DeleteByConditionVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteByConditionVo struct{}"
	}

	return strings.Join([]string{"DeleteByConditionVo", string(data)}, " ")
}
