package nvidia

import (
	"fmt"
	"os/exec"
)

type NVIDIA struct {
	Model string
}

func (nvidia *NVIDIA) Run(prompt string) (string, error) {
	cmd := exec.Command("python3", "nvidia.py", prompt, nvidia.Model)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("error: %v, output: %s", err, string(output))
	}
	return string(output), nil
}
