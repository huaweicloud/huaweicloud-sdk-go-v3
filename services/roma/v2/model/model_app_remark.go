package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AppRemark 应用描述
type AppRemark struct {
}

func (o AppRemark) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppRemark struct{}"
	}

	return strings.Join([]string{"AppRemark", string(data)}, " ")
}
