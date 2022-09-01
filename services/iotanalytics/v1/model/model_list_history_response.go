package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListHistoryResponse struct {

	// 时间序列
	Timestamps *[]string `json:"timestamps,omitempty" xml:"timestamps"`

	// 查询设备的属性值
	Properties     *[]HistoryValues `json:"properties,omitempty" xml:"properties"`
	HttpStatusCode int              `json:"-"`
}

func (o ListHistoryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHistoryResponse struct{}"
	}

	return strings.Join([]string{"ListHistoryResponse", string(data)}, " ")
}
