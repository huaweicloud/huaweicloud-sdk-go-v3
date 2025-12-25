package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportCollectorParserResponse Response Object
type ImportCollectorParserResponse struct {

	// 失败的数量
	Fail *int32 `json:"fail,omitempty"`

	// 解析器数组
	Result *[]ImportParserVo `json:"result,omitempty"`

	// 成功的数量
	Success        *int32 `json:"success,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ImportCollectorParserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportCollectorParserResponse struct{}"
	}

	return strings.Join([]string{"ImportCollectorParserResponse", string(data)}, " ")
}
