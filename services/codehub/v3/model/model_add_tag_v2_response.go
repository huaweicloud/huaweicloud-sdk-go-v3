package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddTagV2Response Response Object
type AddTagV2Response struct {
	Error *Error `json:"error,omitempty"`

	Result *AddTagsResponse `json:"result,omitempty"`

	// 响应状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AddTagV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddTagV2Response struct{}"
	}

	return strings.Join([]string{"AddTagV2Response", string(data)}, " ")
}
