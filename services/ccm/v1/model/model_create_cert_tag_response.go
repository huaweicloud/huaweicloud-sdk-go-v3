package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCertTagResponse Response Object
type CreateCertTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateCertTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCertTagResponse struct{}"
	}

	return strings.Join([]string{"CreateCertTagResponse", string(data)}, " ")
}
