package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateResourceConfigResponse Response Object
type CreateResourceConfigResponse struct {

	// 数据集列表，包含多个数据集对象
	Datasets       *[]DatasetItem `json:"datasets,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o CreateResourceConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResourceConfigResponse struct{}"
	}

	return strings.Join([]string{"CreateResourceConfigResponse", string(data)}, " ")
}
