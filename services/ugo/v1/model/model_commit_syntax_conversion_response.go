package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CommitSyntaxConversionResponse Response Object
type CommitSyntaxConversionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CommitSyntaxConversionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommitSyntaxConversionResponse struct{}"
	}

	return strings.Join([]string{"CommitSyntaxConversionResponse", string(data)}, " ")
}
