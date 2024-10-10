package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BackupInfo 备份文件信息。
type BackupInfo struct {

	// 备份文件来源：（全部大写） - OBS：存放在OBS中的备份文件。 - RDS：RDS实例的备份文件。
	FileSource BackupInfoFileSource `json:"file_source"`

	// OBS桶名称，使用OBS桶备份文件恢复时必填。 约束：长度范围为3～63个字符，仅支持小写字母、数字、中划线（-）、点（.），且只能以字母或数字开头和结尾，禁止使用IP地址。 如果选择从OBS桶获取备份文件必填，参考登录[客户端登录](https://support.huaweicloud.com/clientogw-obs/zh-cn_topic_0045829058.html)OBS Browser客户端，首页为OBS桶列表，可以获取到备份文件所在桶的名称。
	BucketName *string `json:"bucket_name,omitempty"`

	// 备份文件信息列表。
	Files []BackupFileInfo `json:"files"`
}

func (o BackupInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupInfo struct{}"
	}

	return strings.Join([]string{"BackupInfo", string(data)}, " ")
}

type BackupInfoFileSource struct {
	value string
}

type BackupInfoFileSourceEnum struct {
	OBS BackupInfoFileSource
	RDS BackupInfoFileSource
}

func GetBackupInfoFileSourceEnum() BackupInfoFileSourceEnum {
	return BackupInfoFileSourceEnum{
		OBS: BackupInfoFileSource{
			value: "OBS",
		},
		RDS: BackupInfoFileSource{
			value: "RDS",
		},
	}
}

func (c BackupInfoFileSource) Value() string {
	return c.value
}

func (c BackupInfoFileSource) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BackupInfoFileSource) UnmarshalJSON(b []byte) error {
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
