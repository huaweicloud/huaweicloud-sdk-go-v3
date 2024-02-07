package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunSurfacePointsReq 表面离散点坐标集生成请求体
type RunSurfacePointsReq struct {

	// x坐标集
	XCoordList []float32 `json:"x_coord_list"`

	// y坐标集
	YCoordList []float32 `json:"y_coord_list"`

	// z坐标集
	ZCoordList []float32 `json:"z_coord_list"`
}

func (o RunSurfacePointsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunSurfacePointsReq struct{}"
	}

	return strings.Join([]string{"RunSurfacePointsReq", string(data)}, " ")
}
