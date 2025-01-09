package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// InstallType 安装方式：   * `QUIET_INSTALL` - 静默安装。     安装命令(静默安装命令)，例: ${FILE_PATH} /S。   * `UNZIP_INSTALL` - 解压安装。     例: unzip ${FILE_PATH}。   * `GUI_INSTALL` - 用户通过GUI界面安装。 install_type为QUIET_INSTALL、UNZIP_INSTALL时install_command非空。 预定义变量将采用以下值: ${FILE_PATH}: 应用安装包在桌面本地的存储路径。
type InstallType struct {
	value string
}

type InstallTypeEnum struct {
	QUIET_INSTALL InstallType
	UNZIP_INSTALL InstallType
	GUI_INSTALL   InstallType
}

func GetInstallTypeEnum() InstallTypeEnum {
	return InstallTypeEnum{
		QUIET_INSTALL: InstallType{
			value: "QUIET_INSTALL",
		},
		UNZIP_INSTALL: InstallType{
			value: "UNZIP_INSTALL",
		},
		GUI_INSTALL: InstallType{
			value: "GUI_INSTALL",
		},
	}
}

func (c InstallType) Value() string {
	return c.value
}

func (c InstallType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *InstallType) UnmarshalJSON(b []byte) error {
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
