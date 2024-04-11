package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListConditionResponse Response Object
type ListConditionResponse struct {

	// 返回的数据信息。
	Data           *interface{} `json:"data,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListConditionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConditionResponse struct{}"
	}

	return strings.Join([]string{"ListConditionResponse", string(data)}, " ")
}
