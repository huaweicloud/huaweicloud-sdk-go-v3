package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckResourceResponse Response Object
type CheckResourceResponse struct {

	// 结果。 - true：表示通过。 - false：表示不通过。
	Result         *bool `json:"result,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o CheckResourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckResourceResponse struct{}"
	}

	return strings.Join([]string{"CheckResourceResponse", string(data)}, " ")
}
