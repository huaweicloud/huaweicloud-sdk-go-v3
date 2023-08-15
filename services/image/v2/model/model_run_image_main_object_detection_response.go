package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunImageMainObjectDetectionResponse Response Object
type RunImageMainObjectDetectionResponse struct {

	// 主体列表集合。
	Result         *[]ImageMainObjectDetectionInstance `json:"result,omitempty"`
	HttpStatusCode int                                 `json:"-"`
}

func (o RunImageMainObjectDetectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunImageMainObjectDetectionResponse struct{}"
	}

	return strings.Join([]string{"RunImageMainObjectDetectionResponse", string(data)}, " ")
}
