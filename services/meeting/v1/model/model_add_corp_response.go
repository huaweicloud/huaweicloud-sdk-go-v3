package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type AddCorpResponse struct {
	// 返回结果

	Value          *string `json:"value,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AddCorpResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddCorpResponse struct{}"
	}

	return strings.Join([]string{"AddCorpResponse", string(data)}, " ")
}
