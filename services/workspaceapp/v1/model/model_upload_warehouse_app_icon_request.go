package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UploadWarehouseAppIconRequest Request Object
type UploadWarehouseAppIconRequest struct {
	Body *UploadWarehouseAppIconRequestBody `json:"body,omitempty" type:"multipart"`
}

func (o UploadWarehouseAppIconRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadWarehouseAppIconRequest struct{}"
	}

	return strings.Join([]string{"UploadWarehouseAppIconRequest", string(data)}, " ")
}
