package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ParamGroupDiffRequestBody struct {

	// 需要进行比较的参数组模板ID。
	SourceId string `json:"source_id"`

	// 需要进行比较的参数组模板ID，需要与源参数组模板的部署形态相同才可比较。
	TargetId string `json:"target_id"`
}

func (o ParamGroupDiffRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ParamGroupDiffRequestBody struct{}"
	}

	return strings.Join([]string{"ParamGroupDiffRequestBody", string(data)}, " ")
}
