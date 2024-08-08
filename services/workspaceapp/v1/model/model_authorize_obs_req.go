package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AuthorizeObsReq 获取ak/sk时候传入的文件名称（带后缀）。
type AuthorizeObsReq struct {

	// 应用名称,名称需满足如下规则: 1. 文件名前缀由可见字符和空格组成，且不能为全空格。 2. 长度范围1~255个字符。 3. 结尾必须是`.msi`或者`.exe`或者`.zip`或者`.rar`。
	AppFileName string `json:"app_file_name"`
}

func (o AuthorizeObsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuthorizeObsReq struct{}"
	}

	return strings.Join([]string{"AuthorizeObsReq", string(data)}, " ")
}
