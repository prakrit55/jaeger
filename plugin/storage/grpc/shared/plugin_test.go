// Copyright (c) 2024 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package shared

import (
	"context"
	"testing"

	"github.com/jaegertracing/jaeger/storage/dependencystore"
	dependencyStoreMocks "github.com/jaegertracing/jaeger/storage/dependencystore/mocks"
	"github.com/jaegertracing/jaeger/storage/spanstore"
	spanStoreMocks "github.com/jaegertracing/jaeger/storage/spanstore/mocks"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

type mockStorageGRPCPlugin struct {
	spanReader *spanStoreMocks.Reader
	spanWriter *spanStoreMocks.Writer
	depsReader *dependencyStoreMocks.Reader
}

type mockArchiveStoragePlugin struct {
	archiveReader *spanStoreMocks.Reader
	archiveWriter *spanStoreMocks.Writer
}

type mockStreamingSpanWriterPlugin struct {
	streamWriter *spanStoreMocks.Writer
}

func (plugin *mockArchiveStoragePlugin) ArchiveSpanReader() spanstore.Reader {
	return plugin.archiveReader
}

func (plugin *mockArchiveStoragePlugin) ArchiveSpanWriter() spanstore.Writer {
	return plugin.archiveWriter
}

func (plugin *mockStorageGRPCPlugin) SpanReader() spanstore.Reader {
	return plugin.spanReader
}

func (plugin *mockStorageGRPCPlugin) SpanWriter() spanstore.Writer {
	return plugin.spanWriter
}

func (plugin *mockStorageGRPCPlugin) DependencyReader() dependencystore.Reader {
	return plugin.depsReader
}

func (plugin *mockStreamingSpanWriterPlugin) StreamingSpanWriter() spanstore.Writer {
	return plugin.streamWriter
}

func TestStorageGRPCPlugin_RegisterHandlers(t *testing.T) {

	plugin := StorageGRPCPlugin{
		Impl:        &mockStorageGRPCPlugin{},
		ArchiveImpl: &mockArchiveStoragePlugin{},
		StreamImpl:  &mockStreamingSpanWriterPlugin{},
	}

	server := grpc.NewServer()

	err := plugin.RegisterHandlers(server)
	assert.NoError(t, err)
}

func TestStorageGRPCPlugin_GRPCServer(t *testing.T) {

	plugin := &StorageGRPCPlugin{
		Impl:        &mockStoragePlugin{},
		ArchiveImpl: &mockArchiveStoragePlugin{},
		StreamImpl:  &mockStreamingSpanWriterPlugin{},
	}

	server := grpc.NewServer()

	err := plugin.GRPCServer(nil, server)
	assert.NoError(t, err)
}

func TestStorageGRPCPlugin_GRPCClient(t *testing.T) {

	clientConn := &grpc.ClientConn{}

	plugin := &StorageGRPCPlugin{}

	_, err := plugin.GRPCClient(context.Background(), nil, clientConn)
	assert.NoError(t, err)
}
