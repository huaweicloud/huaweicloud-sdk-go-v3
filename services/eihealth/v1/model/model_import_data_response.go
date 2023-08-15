package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportDataResponse Response Object
type ImportDataResponse struct {

	// 数据作业ID
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ImportDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportDataResponse struct{}"
	}

	return strings.Join([]string{"ImportDataResponse", string(data)}, " ")
}
