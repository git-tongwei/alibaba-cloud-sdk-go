package cms

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

func (client *Client) QueryMetricList(request *QueryMetricListRequest) (response *QueryMetricListResponse, err error) {
	response = CreateQueryMetricListResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) QueryMetricListWithChan(request *QueryMetricListRequest) (<-chan *QueryMetricListResponse, <-chan error) {
	responseChan := make(chan *QueryMetricListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryMetricList(request)
		responseChan <- response
		errChan <- err
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

func (client *Client) QueryMetricListWithCallback(request *QueryMetricListRequest, callback func(response *QueryMetricListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryMetricListResponse
		var err error
		defer close(result)
		response, err = client.QueryMetricList(request)
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

type QueryMetricListRequest struct {
	*requests.RpcRequest
	EndTime         string `position:"Query" name:"EndTime"`
	StartTime       string `position:"Query" name:"StartTime"`
	Cursor          string `position:"Query" name:"Cursor"`
	Express         string `position:"Query" name:"Express"`
	Period          string `position:"Query" name:"Period"`
	Project         string `position:"Query" name:"Project"`
	Page            string `position:"Query" name:"Page"`
	Metric          string `position:"Query" name:"Metric"`
	Length          string `position:"Query" name:"Length"`
	Dimensions      string `position:"Query" name:"Dimensions"`
	ResourceOwnerId string `position:"Query" name:"ResourceOwnerId"`
	CallbyCmsOwner  string `position:"Query" name:"callby_cms_owner"`
}

type QueryMetricListResponse struct {
	*responses.BaseResponse
	Code       string `json:"Code" xml:"Code"`
	Message    string `json:"Message" xml:"Message"`
	Success    bool   `json:"Success" xml:"Success"`
	RequestId  string `json:"RequestId" xml:"RequestId"`
	Cursor     string `json:"Cursor" xml:"Cursor"`
	Datapoints string `json:"Datapoints" xml:"Datapoints"`
	Period     string `json:"Period" xml:"Period"`
}

func CreateQueryMetricListRequest() (request *QueryMetricListRequest) {
	request = &QueryMetricListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2017-03-01", "QueryMetricList", "", "")
	return
}

func CreateQueryMetricListResponse() (response *QueryMetricListResponse) {
	response = &QueryMetricListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
