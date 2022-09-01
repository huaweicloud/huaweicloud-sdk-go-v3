package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UploadProductResponse struct {

	// 导入成功的产品数
	SuccNum *int32 `json:"succ_num,omitempty" xml:"succ_num"`

	// 导入失败的产品数
	FailNum *int32 `json:"fail_num,omitempty" xml:"fail_num"`

	// 导入失败的产品名称列表
	FailObjectsIds *[]string `json:"fail_objects_ids,omitempty" xml:"fail_objects_ids"`
	HttpStatusCode int       `json:"-"`
}

func (o UploadProductResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadProductResponse struct{}"
	}

	return strings.Join([]string{"UploadProductResponse", string(data)}, " ")
}
