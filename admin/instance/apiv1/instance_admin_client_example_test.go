// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package instance_test

import (
	"context"

	iampb "cloud.google.com/go/iam/apiv1/iampb"
	instancepb "cloud.google.com/go/spanner/admin/instance/apiv1/instancepb"
	instance "github.com/egonelbre/spanner/admin/instance/apiv1"
	"google.golang.org/api/iterator"
)

func ExampleNewInstanceAdminClient() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := instance.NewInstanceAdminClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleNewInstanceAdminRESTClient() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := instance.NewInstanceAdminRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleInstanceAdminClient_CreateInstance() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := instance.NewInstanceAdminClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &instancepb.CreateInstanceRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/spanner/admin/instance/apiv1/instancepb#CreateInstanceRequest.
	}
	op, err := c.CreateInstance(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleInstanceAdminClient_CreateInstanceConfig() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := instance.NewInstanceAdminClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &instancepb.CreateInstanceConfigRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/spanner/admin/instance/apiv1/instancepb#CreateInstanceConfigRequest.
	}
	op, err := c.CreateInstanceConfig(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleInstanceAdminClient_CreateInstancePartition() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := instance.NewInstanceAdminClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &instancepb.CreateInstancePartitionRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/spanner/admin/instance/apiv1/instancepb#CreateInstancePartitionRequest.
	}
	op, err := c.CreateInstancePartition(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleInstanceAdminClient_DeleteInstance() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := instance.NewInstanceAdminClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &instancepb.DeleteInstanceRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/spanner/admin/instance/apiv1/instancepb#DeleteInstanceRequest.
	}
	err = c.DeleteInstance(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleInstanceAdminClient_DeleteInstanceConfig() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := instance.NewInstanceAdminClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &instancepb.DeleteInstanceConfigRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/spanner/admin/instance/apiv1/instancepb#DeleteInstanceConfigRequest.
	}
	err = c.DeleteInstanceConfig(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleInstanceAdminClient_DeleteInstancePartition() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := instance.NewInstanceAdminClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &instancepb.DeleteInstancePartitionRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/spanner/admin/instance/apiv1/instancepb#DeleteInstancePartitionRequest.
	}
	err = c.DeleteInstancePartition(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleInstanceAdminClient_GetIamPolicy() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := instance.NewInstanceAdminClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &iampb.GetIamPolicyRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/iam/apiv1/iampb#GetIamPolicyRequest.
	}
	resp, err := c.GetIamPolicy(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleInstanceAdminClient_GetInstance() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := instance.NewInstanceAdminClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &instancepb.GetInstanceRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/spanner/admin/instance/apiv1/instancepb#GetInstanceRequest.
	}
	resp, err := c.GetInstance(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleInstanceAdminClient_GetInstanceConfig() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := instance.NewInstanceAdminClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &instancepb.GetInstanceConfigRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/spanner/admin/instance/apiv1/instancepb#GetInstanceConfigRequest.
	}
	resp, err := c.GetInstanceConfig(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleInstanceAdminClient_GetInstancePartition() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := instance.NewInstanceAdminClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &instancepb.GetInstancePartitionRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/spanner/admin/instance/apiv1/instancepb#GetInstancePartitionRequest.
	}
	resp, err := c.GetInstancePartition(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleInstanceAdminClient_ListInstanceConfigOperations() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := instance.NewInstanceAdminClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &instancepb.ListInstanceConfigOperationsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/spanner/admin/instance/apiv1/instancepb#ListInstanceConfigOperationsRequest.
	}
	it := c.ListInstanceConfigOperations(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp

		// If you need to access the underlying RPC response,
		// you can do so by casting the `Response` as below.
		// Otherwise, remove this line. Only populated after
		// first call to Next(). Not safe for concurrent access.
		_ = it.Response.(*instancepb.ListInstanceConfigOperationsResponse)
	}
}

func ExampleInstanceAdminClient_ListInstanceConfigs() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := instance.NewInstanceAdminClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &instancepb.ListInstanceConfigsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/spanner/admin/instance/apiv1/instancepb#ListInstanceConfigsRequest.
	}
	it := c.ListInstanceConfigs(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp

		// If you need to access the underlying RPC response,
		// you can do so by casting the `Response` as below.
		// Otherwise, remove this line. Only populated after
		// first call to Next(). Not safe for concurrent access.
		_ = it.Response.(*instancepb.ListInstanceConfigsResponse)
	}
}

func ExampleInstanceAdminClient_ListInstancePartitionOperations() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := instance.NewInstanceAdminClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &instancepb.ListInstancePartitionOperationsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/spanner/admin/instance/apiv1/instancepb#ListInstancePartitionOperationsRequest.
	}
	it := c.ListInstancePartitionOperations(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp

		// If you need to access the underlying RPC response,
		// you can do so by casting the `Response` as below.
		// Otherwise, remove this line. Only populated after
		// first call to Next(). Not safe for concurrent access.
		_ = it.Response.(*instancepb.ListInstancePartitionOperationsResponse)
	}
}

func ExampleInstanceAdminClient_ListInstancePartitions() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := instance.NewInstanceAdminClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &instancepb.ListInstancePartitionsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/spanner/admin/instance/apiv1/instancepb#ListInstancePartitionsRequest.
	}
	it := c.ListInstancePartitions(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp

		// If you need to access the underlying RPC response,
		// you can do so by casting the `Response` as below.
		// Otherwise, remove this line. Only populated after
		// first call to Next(). Not safe for concurrent access.
		_ = it.Response.(*instancepb.ListInstancePartitionsResponse)
	}
}

func ExampleInstanceAdminClient_ListInstances() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := instance.NewInstanceAdminClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &instancepb.ListInstancesRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/spanner/admin/instance/apiv1/instancepb#ListInstancesRequest.
	}
	it := c.ListInstances(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp

		// If you need to access the underlying RPC response,
		// you can do so by casting the `Response` as below.
		// Otherwise, remove this line. Only populated after
		// first call to Next(). Not safe for concurrent access.
		_ = it.Response.(*instancepb.ListInstancesResponse)
	}
}

func ExampleInstanceAdminClient_SetIamPolicy() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := instance.NewInstanceAdminClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &iampb.SetIamPolicyRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/iam/apiv1/iampb#SetIamPolicyRequest.
	}
	resp, err := c.SetIamPolicy(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleInstanceAdminClient_TestIamPermissions() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := instance.NewInstanceAdminClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &iampb.TestIamPermissionsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/iam/apiv1/iampb#TestIamPermissionsRequest.
	}
	resp, err := c.TestIamPermissions(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleInstanceAdminClient_UpdateInstance() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := instance.NewInstanceAdminClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &instancepb.UpdateInstanceRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/spanner/admin/instance/apiv1/instancepb#UpdateInstanceRequest.
	}
	op, err := c.UpdateInstance(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleInstanceAdminClient_UpdateInstanceConfig() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := instance.NewInstanceAdminClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &instancepb.UpdateInstanceConfigRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/spanner/admin/instance/apiv1/instancepb#UpdateInstanceConfigRequest.
	}
	op, err := c.UpdateInstanceConfig(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleInstanceAdminClient_UpdateInstancePartition() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := instance.NewInstanceAdminClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &instancepb.UpdateInstancePartitionRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/spanner/admin/instance/apiv1/instancepb#UpdateInstancePartitionRequest.
	}
	op, err := c.UpdateInstancePartition(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}
