package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListReposDetailsResponse struct {
	Body *[]ShowReposResp `json:"body,omitempty" xml:"body"`

	ContentRange   *string `json:"Content-Range,omitempty" xml:"Content-Range"`
	HttpStatusCode int     `json:"-"`
}

func (o ListReposDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListReposDetailsResponse struct{}"
	}

	return strings.Join([]string{"ListReposDetailsResponse", string(data)}, " ")
}
