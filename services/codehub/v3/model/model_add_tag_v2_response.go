package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type AddTagV2Response struct {
	Error *Error `json:"error,omitempty" xml:"error"`

	Result *AddTagsResponse `json:"result,omitempty" xml:"result"`

	// 响应状态
	Status         *string `json:"status,omitempty" xml:"status"`
	HttpStatusCode int     `json:"-"`
}

func (o AddTagV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddTagV2Response struct{}"
	}

	return strings.Join([]string{"AddTagV2Response", string(data)}, " ")
}
