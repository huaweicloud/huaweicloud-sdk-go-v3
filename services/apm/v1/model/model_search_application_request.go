package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchApplicationRequest Request Object
type SearchApplicationRequest struct {

	// 应用id。
	XBusinessId int64 `json:"x-business-id"`

	Body *AppSearchParam `json:"body,omitempty"`
}

func (o SearchApplicationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchApplicationRequest struct{}"
	}

	return strings.Join([]string{"SearchApplicationRequest", string(data)}, " ")
}
