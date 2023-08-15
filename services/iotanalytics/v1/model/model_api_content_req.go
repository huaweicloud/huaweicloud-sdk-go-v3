package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ApiContentReq Api数据源请求内容
type ApiContentReq struct {

	// api数据源名称
	Name string `json:"name"`
}

func (o ApiContentReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiContentReq struct{}"
	}

	return strings.Join([]string{"ApiContentReq", string(data)}, " ")
}
