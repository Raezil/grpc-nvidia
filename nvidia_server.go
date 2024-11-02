package nvidia

import (
	"context"
)

type NvidiaServer struct {
	UnimplementedNVIDIAServer
}

func (s *NvidiaServer) Generate(ctx context.Context, req *NVIDIARequest) (*NVIDIAResponse, error) {
	client := NVIDIA{
		Model: req.Model,
	}
	result, err := client.Run(req.Prompt)

	if err != nil {
		return nil, err
	}
	return &NVIDIAResponse{
		Result: result,
	}, nil
}
