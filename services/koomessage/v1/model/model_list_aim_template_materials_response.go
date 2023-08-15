package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAimTemplateMaterialsResponse Response Object
type ListAimTemplateMaterialsResponse struct {

	// 状态码。
	Status *string `json:"status,omitempty"`

	// 响应信息。
	Message *string `json:"message,omitempty"`

	Data           *ListAimTemplateMaterialsResponseMode `json:"data,omitempty"`
	HttpStatusCode int                                   `json:"-"`
}

func (o ListAimTemplateMaterialsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAimTemplateMaterialsResponse struct{}"
	}

	return strings.Join([]string{"ListAimTemplateMaterialsResponse", string(data)}, " ")
}
