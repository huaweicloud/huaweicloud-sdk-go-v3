package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QueryLtsStructLogResponseBodyNew 此参数在请求实体中，采用json字符串格式。
type QueryLtsStructLogResponseBodyNew struct {
}

func (o QueryLtsStructLogResponseBodyNew) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryLtsStructLogResponseBodyNew struct{}"
	}

	return strings.Join([]string{"QueryLtsStructLogResponseBodyNew", string(data)}, " ")
}
