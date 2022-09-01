package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type RecordObsFileAddr struct {

	// OBS Bucket所在RegionID
	Location RecordObsFileAddrLocation `json:"location" xml:"location"`

	// OBS Bucket所在Region的项目ID
	ProjectId string `json:"project_id" xml:"project_id"`

	// OBS的bucket名称
	Bucket string `json:"bucket" xml:"bucket"`

	// OBS对象路径，遵守OBS Object定义。如果为空则保存到根目录
	Object *string `json:"object,omitempty" xml:"object"`
}

func (o RecordObsFileAddr) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecordObsFileAddr struct{}"
	}

	return strings.Join([]string{"RecordObsFileAddr", string(data)}, " ")
}

type RecordObsFileAddrLocation struct {
	value string
}

type RecordObsFileAddrLocationEnum struct {
	CN_NORTH_4 RecordObsFileAddrLocation
}

func GetRecordObsFileAddrLocationEnum() RecordObsFileAddrLocationEnum {
	return RecordObsFileAddrLocationEnum{
		CN_NORTH_4: RecordObsFileAddrLocation{
			value: "cn-north-4",
		},
	}
}

func (c RecordObsFileAddrLocation) Value() string {
	return c.value
}

func (c RecordObsFileAddrLocation) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RecordObsFileAddrLocation) UnmarshalJSON(b []byte) error {
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
