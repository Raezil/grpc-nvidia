package nvidia

import (
	context "context"
	"fmt"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestNvidia(t *testing.T) {
	// Create a mock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock of the AuthClient
	mockAuthClient := NewMockNVIDIAClient(ctrl)

	// Define the input and expected output for registration
	regReq := &NVIDIARequest{
		Prompt: "Tell me more about golang",
		Model:  "nemotron-instruct",
	}
	regReply := fmt.Sprintf("Go is a statically typed, compiled high-level general purpose programming language. It was designed at Google in 2009 by Robert Griesemer, Rob Pike, and Ken Thompson. It is syntactically similar to C, but also has memory safety, garbage collection, structural typing, and CSP-style concurrency.")
	expectedRegRes := &NVIDIAResponse{Result: regReply}

	// Set up the mock to expect a Register call and return the expected response
	mockAuthClient.EXPECT().Generate(gomock.Any(), gomock.Eq(regReq), gomock.Any()).
		Return(expectedRegRes, nil)

	// Call the Register method
	regRes, regErr := mockAuthClient.Generate(context.Background(), regReq)

	// Validate the registration response
	assert.NoError(t, regErr)
	assert.Equal(t, expectedRegRes, regRes)
}
