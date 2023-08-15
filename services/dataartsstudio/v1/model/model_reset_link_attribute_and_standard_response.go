package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResetLinkAttributeAndStandardResponse Response Object
type ResetLinkAttributeAndStandardResponse struct {

	// 返回的数据信息
	Data           *interface{} `json:"data,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ResetLinkAttributeAndStandardResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetLinkAttributeAndStandardResponse struct{}"
	}

	return strings.Join([]string{"ResetLinkAttributeAndStandardResponse", string(data)}, " ")
}
