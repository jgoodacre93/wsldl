package run

import (
	"runtime"
	"strings"
	"testing"

	"github.com/yuk7/wsldl/lib/wsllib"
)

func TestExecWindowsTerminal_NonWindows_ReturnsDisplayError(t *testing.T) {
	t.Parallel()
	if runtime.GOOS == "windows" {
		t.Skip("skip on windows to avoid launching real terminal process")
	}

	reg := wsllib.MockWslReg{
		GetProfileFromNameFunc: func(name string) (wsllib.Profile, error) {
			return wsllib.Profile{
				DistributionName: "ArchLinux",
			}, nil
		},
	}

	err := ExecWindowsTerminal(reg, "arch")
	de := assertDisplayError(t, err)
	if !strings.Contains(de.Error(), "unsupported platform") {
		t.Fatalf("error = %q, want to contain %q", de.Error(), "unsupported platform")
	}
}
