package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckStarRocksResourceResponse Response Object
type CheckStarRocksResourceResponse struct {

	// 检查结果。
	Result         *bool `json:"result,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o CheckStarRocksResourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckStarRocksResourceResponse struct{}"
	}

	return strings.Join([]string{"CheckStarRocksResourceResponse", string(data)}, " ")
}
