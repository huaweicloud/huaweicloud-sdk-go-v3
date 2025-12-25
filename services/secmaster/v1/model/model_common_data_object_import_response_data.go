package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CommonDataObjectImportResponseData 数据对象导入结果返回数据
type CommonDataObjectImportResponseData struct {

	// 导入的数据对象的id列表
	Ids *[]string `json:"ids,omitempty"`
}

func (o CommonDataObjectImportResponseData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommonDataObjectImportResponseData struct{}"
	}

	return strings.Join([]string{"CommonDataObjectImportResponseData", string(data)}, " ")
}
