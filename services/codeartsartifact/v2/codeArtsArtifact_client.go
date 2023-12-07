package v2

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/codeartsartifact/v2/model"
)

type CodeArtsArtifactClient struct {
	HcClient *http_client.HcHttpClient
}

func NewCodeArtsArtifactClient(hcClient *http_client.HcHttpClient) *CodeArtsArtifactClient {
	return &CodeArtsArtifactClient{HcClient: hcClient}
}

func CodeArtsArtifactClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// ShowProjectReleaseFiles 获取项目下文件版本信息列表
//
// 获取项目下文件版本信息列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsArtifactClient) ShowProjectReleaseFiles(request *model.ShowProjectReleaseFilesRequest) (*model.ShowProjectReleaseFilesResponse, error) {
	requestDef := GenReqDefForShowProjectReleaseFiles()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowProjectReleaseFilesResponse), nil
	}
}

// ShowProjectReleaseFilesInvoker 获取项目下文件版本信息列表
func (c *CodeArtsArtifactClient) ShowProjectReleaseFilesInvoker(request *model.ShowProjectReleaseFilesRequest) *ShowProjectReleaseFilesInvoker {
	requestDef := GenReqDefForShowProjectReleaseFiles()
	return &ShowProjectReleaseFilesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// ShowReleaseProjectFiles 获取项目下文件版本信息列表
//
// 获取项目下文件版本信息列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsArtifactClient) ShowReleaseProjectFiles(request *model.ShowReleaseProjectFilesRequest) (*model.ShowReleaseProjectFilesResponse, error) {
	requestDef := GenReqDefForShowReleaseProjectFiles()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowReleaseProjectFilesResponse), nil
	}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// ShowReleaseProjectFilesInvoker 获取项目下文件版本信息列表
func (c *CodeArtsArtifactClient) ShowReleaseProjectFilesInvoker(request *model.ShowReleaseProjectFilesRequest) *ShowReleaseProjectFilesInvoker {
	requestDef := GenReqDefForShowReleaseProjectFiles()
	return &ShowReleaseProjectFilesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
