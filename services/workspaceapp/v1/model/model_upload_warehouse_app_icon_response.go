package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UploadWarehouseAppIconResponse Response Object
type UploadWarehouseAppIconResponse struct {

	// 图标文件在obs桶经过cdn加速以后的地址。
	AppiconStorePath *string `json:"appicon_store_path,omitempty"`
	HttpStatusCode   int     `json:"-"`
}

func (o UploadWarehouseAppIconResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadWarehouseAppIconResponse struct{}"
	}

	return strings.Join([]string{"UploadWarehouseAppIconResponse", string(data)}, " ")
}
