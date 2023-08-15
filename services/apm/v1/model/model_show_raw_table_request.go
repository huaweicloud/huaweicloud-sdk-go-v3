package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRawTableRequest Request Object
type ShowRawTableRequest struct {

	// 应用id。
	XBusinessId int64 `json:"x-business-id"`

	Body *RawTableParam `json:"body,omitempty"`
}

func (o ShowRawTableRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRawTableRequest struct{}"
	}

	return strings.Join([]string{"ShowRawTableRequest", string(data)}, " ")
}
