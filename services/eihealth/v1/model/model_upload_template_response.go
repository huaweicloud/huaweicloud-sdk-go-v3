package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UploadTemplateResponse Response Object
type UploadTemplateResponse struct {

	// 模板id
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UploadTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadTemplateResponse struct{}"
	}

	return strings.Join([]string{"UploadTemplateResponse", string(data)}, " ")
}
