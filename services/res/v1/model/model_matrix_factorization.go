package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 矩阵分解参数配置
type MatrixFactorization struct {

	// 隐向量维度。
	ImplicitVectorRank int32 `json:"implicit_vector_rank"`

	// 优化正则化系数。
	RegularParam float64 `json:"regular_param"`

	// 迭代次数。
	MaxIteratorNum int32 `json:"max_iterator_num"`
}

func (o MatrixFactorization) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MatrixFactorization struct{}"
	}

	return strings.Join([]string{"MatrixFactorization", string(data)}, " ")
}
