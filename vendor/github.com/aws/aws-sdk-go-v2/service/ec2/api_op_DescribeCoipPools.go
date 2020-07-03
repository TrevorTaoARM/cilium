// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeCoipPoolsInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The filters. The following are the possible values:
	//
	//    * coip-pool.pool-id
	//
	//    * coip-pool.local-gateway-route-table-id
	Filters []Filter `locationName:"Filter" locationNameList:"Filter" type:"list"`

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int64 `min:"5" type:"integer"`

	// The token for the next page of results.
	NextToken *string `type:"string"`

	// The IDs of the address pools.
	PoolIds []string `locationName:"PoolId" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s DescribeCoipPoolsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeCoipPoolsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeCoipPoolsInput"}
	if s.MaxResults != nil && *s.MaxResults < 5 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 5))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeCoipPoolsOutput struct {
	_ struct{} `type:"structure"`

	// Information about the address pools.
	CoipPools []CoipPool `locationName:"coipPoolSet" locationNameList:"item" type:"list"`

	// The token to use to retrieve the next page of results. This value is null
	// when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s DescribeCoipPoolsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeCoipPools = "DescribeCoipPools"

// DescribeCoipPoolsRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Describes the specified customer-owned address pools or all of your customer-owned
// address pools.
//
//    // Example sending a request using DescribeCoipPoolsRequest.
//    req := client.DescribeCoipPoolsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeCoipPools
func (c *Client) DescribeCoipPoolsRequest(input *DescribeCoipPoolsInput) DescribeCoipPoolsRequest {
	op := &aws.Operation{
		Name:       opDescribeCoipPools,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeCoipPoolsInput{}
	}

	req := c.newRequest(op, input, &DescribeCoipPoolsOutput{})

	return DescribeCoipPoolsRequest{Request: req, Input: input, Copy: c.DescribeCoipPoolsRequest}
}

// DescribeCoipPoolsRequest is the request type for the
// DescribeCoipPools API operation.
type DescribeCoipPoolsRequest struct {
	*aws.Request
	Input *DescribeCoipPoolsInput
	Copy  func(*DescribeCoipPoolsInput) DescribeCoipPoolsRequest
}

// Send marshals and sends the DescribeCoipPools API request.
func (r DescribeCoipPoolsRequest) Send(ctx context.Context) (*DescribeCoipPoolsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeCoipPoolsResponse{
		DescribeCoipPoolsOutput: r.Request.Data.(*DescribeCoipPoolsOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeCoipPoolsRequestPaginator returns a paginator for DescribeCoipPools.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeCoipPoolsRequest(input)
//   p := ec2.NewDescribeCoipPoolsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeCoipPoolsPaginator(req DescribeCoipPoolsRequest) DescribeCoipPoolsPaginator {
	return DescribeCoipPoolsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeCoipPoolsInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// DescribeCoipPoolsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeCoipPoolsPaginator struct {
	aws.Pager
}

func (p *DescribeCoipPoolsPaginator) CurrentPage() *DescribeCoipPoolsOutput {
	return p.Pager.CurrentPage().(*DescribeCoipPoolsOutput)
}

// DescribeCoipPoolsResponse is the response type for the
// DescribeCoipPools API operation.
type DescribeCoipPoolsResponse struct {
	*DescribeCoipPoolsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeCoipPools request.
func (r *DescribeCoipPoolsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
