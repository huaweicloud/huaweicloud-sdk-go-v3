package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 添加图像数据的相关信息，不同服务类型返回信息不同，具体可参见服务类型说明。
type AddDataRestInfoImageInfo struct {

	// 添加的人脸数量。
	FaceNum float32 `json:"face_num,omitempty"`

	// 添加的主体列表。
	Objects *[]AddDataRestInfoImageInfoObjects `json:"objects,omitempty"`
}

func (o AddDataRestInfoImageInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddDataRestInfoImageInfo struct{}"
	}

	return strings.Join([]string{"AddDataRestInfoImageInfo", string(data)}, " ")
}
