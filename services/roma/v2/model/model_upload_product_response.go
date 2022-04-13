package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UploadProductResponse struct {
	// 导入成功的产品数

	SuccNum *int32 `json:"succ_num,omitempty"`
	// 导入失败的产品数

	FailNum *int32 `json:"fail_num,omitempty"`
	// 导入失败的产品名称列表

	FailObjectsIds *[]string `json:"fail_objects_ids,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o UploadProductResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadProductResponse struct{}"
	}

	return strings.Join([]string{"UploadProductResponse", string(data)}, " ")
}
