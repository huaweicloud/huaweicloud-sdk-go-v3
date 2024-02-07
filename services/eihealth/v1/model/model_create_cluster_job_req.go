package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateClusterJobReq 创建分子聚类作业请求体
type CreateClusterJobReq struct {

	// 聚类方法,当前仅支持hiq_mc
	Method string `json:"method"`

	// 分子聚类源数据
	File string `json:"file"`

	// 分子聚类输出结果
	OutputDir string `json:"output_dir"`
}

func (o CreateClusterJobReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateClusterJobReq struct{}"
	}

	return strings.Join([]string{"CreateClusterJobReq", string(data)}, " ")
}
