package v2

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cloudartifact/v2/model"
)

type CloudArtifactClient struct {
	HcClient *http_client.HcHttpClient
}

func NewCloudArtifactClient(hcClient *http_client.HcHttpClient) *CloudArtifactClient {
	return &CloudArtifactClient{HcClient: hcClient}
}

func CloudArtifactClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// ShowProjectReleaseFiles 获取项目下文件版本信息列表
//
// 获取项目下文件版本信息列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudArtifactClient) ShowProjectReleaseFiles(request *model.ShowProjectReleaseFilesRequest) (*model.ShowProjectReleaseFilesResponse, error) {
	requestDef := GenReqDefForShowProjectReleaseFiles()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowProjectReleaseFilesResponse), nil
	}
}

// ShowProjectReleaseFilesInvoker 获取项目下文件版本信息列表
func (c *CloudArtifactClient) ShowProjectReleaseFilesInvoker(request *model.ShowProjectReleaseFilesRequest) *ShowProjectReleaseFilesInvoker {
	requestDef := GenReqDefForShowProjectReleaseFiles()
	return &ShowProjectReleaseFilesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowReleaseProjectFiles 获取项目下文件版本信息列表
//
// 获取项目下文件版本信息列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudArtifactClient) ShowReleaseProjectFiles(request *model.ShowReleaseProjectFilesRequest) (*model.ShowReleaseProjectFilesResponse, error) {
	requestDef := GenReqDefForShowReleaseProjectFiles()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowReleaseProjectFilesResponse), nil
	}
}

// ShowReleaseProjectFilesInvoker 获取项目下文件版本信息列表
func (c *CloudArtifactClient) ShowReleaseProjectFilesInvoker(request *model.ShowReleaseProjectFilesRequest) *ShowReleaseProjectFilesInvoker {
	requestDef := GenReqDefForShowReleaseProjectFiles()
	return &ShowReleaseProjectFilesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
