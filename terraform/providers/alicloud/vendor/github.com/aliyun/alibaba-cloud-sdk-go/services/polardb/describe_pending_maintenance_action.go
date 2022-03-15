package polardb

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DescribePendingMaintenanceAction invokes the polardb.DescribePendingMaintenanceAction API synchronously
func (client *Client) DescribePendingMaintenanceAction(request *DescribePendingMaintenanceActionRequest) (response *DescribePendingMaintenanceActionResponse, err error) {
	response = CreateDescribePendingMaintenanceActionResponse()
	err = client.DoAction(request, response)
	return
}

// DescribePendingMaintenanceActionWithChan invokes the polardb.DescribePendingMaintenanceAction API asynchronously
func (client *Client) DescribePendingMaintenanceActionWithChan(request *DescribePendingMaintenanceActionRequest) (<-chan *DescribePendingMaintenanceActionResponse, <-chan error) {
	responseChan := make(chan *DescribePendingMaintenanceActionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribePendingMaintenanceAction(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DescribePendingMaintenanceActionWithCallback invokes the polardb.DescribePendingMaintenanceAction API asynchronously
func (client *Client) DescribePendingMaintenanceActionWithCallback(request *DescribePendingMaintenanceActionRequest, callback func(response *DescribePendingMaintenanceActionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribePendingMaintenanceActionResponse
		var err error
		defer close(result)
		response, err = client.DescribePendingMaintenanceAction(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// DescribePendingMaintenanceActionRequest is the request struct for api DescribePendingMaintenanceAction
type DescribePendingMaintenanceActionRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	IsHistory            requests.Integer `position:"Query" name:"IsHistory"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	TaskType             string           `position:"Query" name:"TaskType"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Region               string           `position:"Query" name:"Region"`
}

// DescribePendingMaintenanceActionResponse is the response struct for api DescribePendingMaintenanceAction
type DescribePendingMaintenanceActionResponse struct {
	*responses.BaseResponse
	TotalRecordCount int         `json:"TotalRecordCount" xml:"TotalRecordCount"`
	RequestId        string      `json:"RequestId" xml:"RequestId"`
	PageSize         int         `json:"PageSize" xml:"PageSize"`
	PageNumber       int         `json:"PageNumber" xml:"PageNumber"`
	Items            []ItemsItem `json:"Items" xml:"Items"`
}

// CreateDescribePendingMaintenanceActionRequest creates a request to invoke DescribePendingMaintenanceAction API
func CreateDescribePendingMaintenanceActionRequest() (request *DescribePendingMaintenanceActionRequest) {
	request = &DescribePendingMaintenanceActionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("polardb", "2017-08-01", "DescribePendingMaintenanceAction", "polardb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribePendingMaintenanceActionResponse creates a response to parse from DescribePendingMaintenanceAction response
func CreateDescribePendingMaintenanceActionResponse() (response *DescribePendingMaintenanceActionResponse) {
	response = &DescribePendingMaintenanceActionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}