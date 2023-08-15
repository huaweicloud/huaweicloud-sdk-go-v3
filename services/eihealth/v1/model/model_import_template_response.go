package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportTemplateResponse Response Object
type ImportTemplateResponse struct {
	Body           *[]ImportTemplateResultRsp `json:"body,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o ImportTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportTemplateResponse struct{}"
	}

	return strings.Join([]string{"ImportTemplateResponse", string(data)}, " ")
}
