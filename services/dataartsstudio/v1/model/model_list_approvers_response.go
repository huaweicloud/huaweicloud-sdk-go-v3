package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListApproversResponse Response Object
type ListApproversResponse struct {

	// 返回的数据信息
	Data           *interface{} `json:"data,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListApproversResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApproversResponse struct{}"
	}

	return strings.Join([]string{"ListApproversResponse", string(data)}, " ")
}
