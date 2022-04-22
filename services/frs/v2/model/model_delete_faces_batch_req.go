package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteFacesBatchReq struct {

	// 过滤条件，参考[filter语法](https://support.huaweicloud.com/api-face/face_02_0014.html)。
	Filter string `json:"filter"`
}

func (o DeleteFacesBatchReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteFacesBatchReq struct{}"
	}

	return strings.Join([]string{"DeleteFacesBatchReq", string(data)}, " ")
}
