package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UploadCsrResponse Response Object
type UploadCsrResponse struct {

	// CSR的ID。
	Id *string `json:"id,omitempty"`

	// 自定义CSR名称。
	Name *string `json:"name,omitempty"`

	// CSR创建时间。
	CreateTime     float32 `json:"create_time,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UploadCsrResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadCsrResponse struct{}"
	}

	return strings.Join([]string{"UploadCsrResponse", string(data)}, " ")
}
