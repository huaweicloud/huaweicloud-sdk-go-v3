package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInferServiceTagsResponse Response Object
type ShowInferServiceTagsResponse struct {

	// **参数解释：** 资源标签返回列表。
	Tags           *[]InferTmsTag `json:"tags,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowInferServiceTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInferServiceTagsResponse struct{}"
	}

	return strings.Join([]string{"ShowInferServiceTagsResponse", string(data)}, " ")
}
