package git

import (
	"context"
	"os/exec"
)

func GetRemoteOrigin(ctx context.Context) ([]byte, error) {
	return exec.CommandContext(ctx, "git", "config", "--get", "remote.origin.url").Output()
}
