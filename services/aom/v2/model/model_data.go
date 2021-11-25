package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Data struct {
	ResultType *string `json:"resultType,omitempty"`

	Result *[]string `json:"result,omitempty"`
}

func (o Data) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Data struct{}"
	}

	return strings.Join([]string{"Data", string(data)}, " ")
}
