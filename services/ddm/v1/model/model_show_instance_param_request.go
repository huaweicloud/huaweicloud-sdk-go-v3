package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowInstanceParamRequest struct {
	// DDM实例ID

	InstanceId string `json:"instance_id"`
	// 分页参数：起始值 [大于等于0] 。

	Offset *int32 `json:"offset,omitempty"`
	// 分页参数：每页多少条 [大于0且小于等于128]。

	Limit *int32 `json:"limit,omitempty"`
	// 语种，默认中文。中文:zh-cn;英文:en-us

	XLanguage *string `json:"X-Language,omitempty"`
}

func (o ShowInstanceParamRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowInstanceParamRequest struct{}"
	}

	return strings.Join([]string{"ShowInstanceParamRequest", string(data)}, " ")
}
