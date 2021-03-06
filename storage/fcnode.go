package storage

import (
	"context"
	"time"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (fc *fcstorage) NodePublishVolume(
	ctx context.Context,
	req *csi.NodePublishVolumeRequest) (
	*csi.NodePublishVolumeResponse, error) {
	return &csi.NodePublishVolumeResponse{}, status.Error(codes.Unimplemented, time.Now().String())

}
func (fc *fcstorage) NodeUnpublishVolume(
	ctx context.Context,
	req *csi.NodeUnpublishVolumeRequest) (
	*csi.NodeUnpublishVolumeResponse, error) {
	return &csi.NodeUnpublishVolumeResponse{}, status.Error(codes.Unimplemented, time.Now().String())
}

func (fc *fcstorage) NodeGetCapabilities(
	ctx context.Context,
	req *csi.NodeGetCapabilitiesRequest) (
	*csi.NodeGetCapabilitiesResponse, error) {
	return &csi.NodeGetCapabilitiesResponse{}, status.Error(codes.Unimplemented, time.Now().String())
}

func (fc *fcstorage) NodeGetInfo(
	ctx context.Context,
	req *csi.NodeGetInfoRequest) (
	*csi.NodeGetInfoResponse, error) {
	return &csi.NodeGetInfoResponse{}, status.Error(codes.Unimplemented, time.Now().String())
}

func (fc *fcstorage) NodeStageVolume(
	ctx context.Context,
	req *csi.NodeStageVolumeRequest) (
	*csi.NodeStageVolumeResponse, error) {
	return &csi.NodeStageVolumeResponse{}, status.Error(codes.Unimplemented, time.Now().String())
}

func (fc *fcstorage) NodeUnstageVolume(
	ctx context.Context,
	req *csi.NodeUnstageVolumeRequest) (
	*csi.NodeUnstageVolumeResponse, error) {
	return &csi.NodeUnstageVolumeResponse{}, status.Error(codes.Unimplemented, time.Now().String())
}
func (fc *fcstorage) NodeGetVolumeStats(
	ctx context.Context, req *csi.NodeGetVolumeStatsRequest) (*csi.NodeGetVolumeStatsResponse, error) {
	return &csi.NodeGetVolumeStatsResponse{}, status.Error(codes.Unimplemented, time.Now().String())

}

func (fc *fcstorage) NodeExpandVolume(ctx context.Context, req *csi.NodeExpandVolumeRequest) (*csi.NodeExpandVolumeResponse, error) {
	return &csi.NodeExpandVolumeResponse{}, status.Error(codes.Unimplemented, time.Now().String())
}