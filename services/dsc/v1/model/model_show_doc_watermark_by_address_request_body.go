package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ShowDocWatermarkByAddressRequestBody struct {

	// 项目所在region的id，如：xx-xx-1。
	RegionId string `json:"region_id"`

	// 待提取水印的文档类型
	DocType ShowDocWatermarkByAddressRequestBodyDocType `json:"doc_type"`

	// 待提取文字暗水印的文档的地址，当前只支持华为云OBS对象，格式为 **obs://bucket/object** ，其中bucket为和当前项目处于同一区域的OBS桶名称，object为对象全路径名。例如：**obs://hwbucket/hwinfo/hw.doc**，其中obs://表示OBS存储，hwbucket为桶名，hwinfo/hw.doc为对象全路径名。
	SrcFile string `json:"src_file"`

	// 解密文件的密码， 最大支持长度256。如果Office文档有读密码或域控的权限密码，请输入读密码，或者有读权限的域控密码。
	FilePassword *string `json:"file_password,omitempty"`
}

func (o ShowDocWatermarkByAddressRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDocWatermarkByAddressRequestBody struct{}"
	}

	return strings.Join([]string{"ShowDocWatermarkByAddressRequestBody", string(data)}, " ")
}

type ShowDocWatermarkByAddressRequestBodyDocType struct {
	value string
}

type ShowDocWatermarkByAddressRequestBodyDocTypeEnum struct {
	WORD  ShowDocWatermarkByAddressRequestBodyDocType
	EXCEL ShowDocWatermarkByAddressRequestBodyDocType
	PDF   ShowDocWatermarkByAddressRequestBodyDocType
	PPT   ShowDocWatermarkByAddressRequestBodyDocType
}

func GetShowDocWatermarkByAddressRequestBodyDocTypeEnum() ShowDocWatermarkByAddressRequestBodyDocTypeEnum {
	return ShowDocWatermarkByAddressRequestBodyDocTypeEnum{
		WORD: ShowDocWatermarkByAddressRequestBodyDocType{
			value: "WORD",
		},
		EXCEL: ShowDocWatermarkByAddressRequestBodyDocType{
			value: "EXCEL",
		},
		PDF: ShowDocWatermarkByAddressRequestBodyDocType{
			value: "PDF",
		},
		PPT: ShowDocWatermarkByAddressRequestBodyDocType{
			value: "PPT",
		},
	}
}

func (c ShowDocWatermarkByAddressRequestBodyDocType) Value() string {
	return c.value
}

func (c ShowDocWatermarkByAddressRequestBodyDocType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDocWatermarkByAddressRequestBodyDocType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
