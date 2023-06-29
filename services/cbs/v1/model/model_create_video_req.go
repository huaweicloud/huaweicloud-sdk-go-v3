package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateVideoReq This is a auto create Body Object
type CreateVideoReq struct {

	// 视频名，长度3~63位。 大小写字母，数字，汉字及部分符号(“_”、“-”、“#”)组成
	Name *string `json:"name,omitempty"`
}

func (o CreateVideoReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVideoReq struct{}"
	}

	return strings.Join([]string{"CreateVideoReq", string(data)}, " ")
}
