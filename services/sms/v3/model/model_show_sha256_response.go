package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSha256Response Response Object
type ShowSha256Response struct {

	// Sha256值
	Value          *string `json:"value,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowSha256Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSha256Response struct{}"
	}

	return strings.Join([]string{"ShowSha256Response", string(data)}, " ")
}
