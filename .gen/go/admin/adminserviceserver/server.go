// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by thriftrw-plugin-yarpc
// @generated

package adminserviceserver

import (
	"context"
	"github.com/uber/cadence/.gen/go/admin"
	"github.com/uber/cadence/.gen/go/shared"
	"go.uber.org/thriftrw/wire"
	"go.uber.org/yarpc/api/transport"
	"go.uber.org/yarpc/encoding/thrift"
)

// Interface is the server-side interface for the AdminService service.
type Interface interface {
	InquiryWorkflowExecution(
		ctx context.Context,
		InquiryRequest *shared.DescribeWorkflowExecutionRequest,
	) (*admin.InquiryWorkflowExecutionResponse, error)
}

// New prepares an implementation of the AdminService service for
// registration.
//
// 	handler := AdminServiceHandler{}
// 	dispatcher.Register(adminserviceserver.New(handler))
func New(impl Interface, opts ...thrift.RegisterOption) []transport.Procedure {
	h := handler{impl}
	service := thrift.Service{
		Name: "AdminService",
		Methods: []thrift.Method{

			thrift.Method{
				Name: "InquiryWorkflowExecution",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.InquiryWorkflowExecution),
				},
				Signature:    "InquiryWorkflowExecution(InquiryRequest *shared.DescribeWorkflowExecutionRequest) (*admin.InquiryWorkflowExecutionResponse)",
				ThriftModule: admin.ThriftModule,
			},
		},
	}

	procedures := make([]transport.Procedure, 0, 1)
	procedures = append(procedures, thrift.BuildProcedures(service, opts...)...)
	return procedures
}

type handler struct{ impl Interface }

func (h handler) InquiryWorkflowExecution(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args admin.AdminService_InquiryWorkflowExecution_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, err
	}

	success, err := h.impl.InquiryWorkflowExecution(ctx, args.InquiryRequest)

	hadError := err != nil
	result, err := admin.AdminService_InquiryWorkflowExecution_Helper.WrapResponse(success, err)

	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Body = result
	}
	return response, err
}
