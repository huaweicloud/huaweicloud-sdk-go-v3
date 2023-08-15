package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListExtensionsRequest Request Object
type ListExtensionsRequest struct {
	Body *ExtensionQueryParamSnake `json:"body,omitempty"`
}

func (o ListExtensionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListExtensionsRequest struct{}"
	}

	return strings.Join([]string{"ListExtensionsRequest", string(data)}, " ")
}
