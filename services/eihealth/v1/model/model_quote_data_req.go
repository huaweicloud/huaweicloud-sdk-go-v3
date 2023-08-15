package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QuoteDataReq 引用数据请求体
type QuoteDataReq struct {

	// 引入项目ID
	QuoteProjectId string `json:"quote_project_id"`

	// 引入路径集
	SubPaths []string `json:"sub_paths"`
}

func (o QuoteDataReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuoteDataReq struct{}"
	}

	return strings.Join([]string{"QuoteDataReq", string(data)}, " ")
}
