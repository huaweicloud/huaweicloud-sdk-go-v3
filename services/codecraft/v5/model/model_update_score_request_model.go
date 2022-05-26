package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdateScoreRequestModel struct {

	// 作品ID，大赛平台提供，可以通过接口[ListCompetitionWorks](https://apiexplorer.developer.huaweicloud.com/apiexplorer/doc?product=CodeCraft&api=ListCompetitionWorks)查询作品ID
	WorksId int32 `json:"works_id"`

	// 作品分数，作品状态为failed时传-1，计算长度时包括小数点，小数点后面最多保留四位
	Score float64 `json:"score"`

	// 作品状态success|failed。判题时，需要对上传作品进行检查，当作品不符合要求时，应该返回failed，并将提示信息通过 message显示出来
	Status UpdateScoreRequestModelStatus `json:"status"`

	// 作品描述信息
	Message *string `json:"message,omitempty"`
}

func (o UpdateScoreRequestModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateScoreRequestModel struct{}"
	}

	return strings.Join([]string{"UpdateScoreRequestModel", string(data)}, " ")
}

type UpdateScoreRequestModelStatus struct {
	value string
}

type UpdateScoreRequestModelStatusEnum struct {
	SUCCESS UpdateScoreRequestModelStatus
	FAILED  UpdateScoreRequestModelStatus
}

func GetUpdateScoreRequestModelStatusEnum() UpdateScoreRequestModelStatusEnum {
	return UpdateScoreRequestModelStatusEnum{
		SUCCESS: UpdateScoreRequestModelStatus{
			value: "success",
		},
		FAILED: UpdateScoreRequestModelStatus{
			value: "failed",
		},
	}
}

func (c UpdateScoreRequestModelStatus) Value() string {
	return c.value
}

func (c UpdateScoreRequestModelStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateScoreRequestModelStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
