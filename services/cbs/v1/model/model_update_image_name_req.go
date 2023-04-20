package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// This is a auto create Body Object
type UpdateImageNameReq struct {

	// 视频名，长度3~63位。 大小写字母，数字，汉字及部分符号(“_”、“-”、“#”)组成
	Name string `json:"name"`
}

func (o UpdateImageNameReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateImageNameReq struct{}"
	}

	return strings.Join([]string{"UpdateImageNameReq", string(data)}, " ")
}
