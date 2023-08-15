package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UploadAimTemplateMaterialResponse Response Object
type UploadAimTemplateMaterialResponse struct {

	// 状态码。
	Status *string `json:"status,omitempty"`

	// 响应信息。
	Message *string `json:"message,omitempty"`

	Data           *UploadAimTemplateMaterialResponseMode `json:"data,omitempty"`
	HttpStatusCode int                                    `json:"-"`
}

func (o UploadAimTemplateMaterialResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadAimTemplateMaterialResponse struct{}"
	}

	return strings.Join([]string{"UploadAimTemplateMaterialResponse", string(data)}, " ")
}
