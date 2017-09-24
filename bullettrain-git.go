package carGit

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/bullettrain-sh/bullettrain-go-core/pkg/ansi"
)

const (
	carPaint       = "black:white"
	gitSymbolPaint = "red:white"
	gitSymbolIcon  = " "
	gitDirtyPaint  = "red:white"
	gitDirtyIcon   = "✘"
	gitCleanPaint  = "green:white"
	gitCleanIcon   = "✔"
)

// Car for Git
type Car struct {
	paint string
	// Current directory
	Pwd string
}

func paintedSymbol() string {
	var symbolIcon string
	if symbolIcon = os.Getenv("BULLETTRAIN_CAR_GIT_ICON"); symbolIcon == "" {
		symbolIcon = gitSymbolIcon
	}

	var symbolPaint string
	if symbolPaint = os.Getenv("BULLETTRAIN_CAR_GIT_ICON_PAINT"); symbolPaint == "" {
		symbolPaint = gitSymbolPaint
	}

	return ansi.Color(symbolIcon, symbolPaint)
}

func paintStatus(pwd string) string {
	var dirtyIcon string
	if dirtyIcon = os.Getenv("BULLETTRAIN_CAR_GIT_DIRTY_ICON"); dirtyIcon == "" {
		dirtyIcon = gitDirtyIcon
	}

	var dirtyPaint string
	if dirtyPaint = os.Getenv("BULLETTRAIN_CAR_GIT_DIRTY_PAINT"); dirtyPaint == "" {
		dirtyPaint = gitDirtyPaint
	}

	var cleanIcon string
	if cleanIcon = os.Getenv("BULLETTRAIN_CAR_GIT_CLEAN_ICON"); cleanIcon == "" {
		cleanIcon = gitCleanIcon
	}

	var cleanPaint string
	if cleanPaint = os.Getenv("BULLETTRAIN_CAR_GIT_CLEAN_PAINT"); cleanPaint == "" {
		cleanPaint = gitCleanPaint
	}

	cmd := exec.Command("git", "-C", pwd, "status", "--porcelain")
	out, _ := cmd.Output()
	if len(out) > 0 {
		return ansi.Color(dirtyIcon, dirtyPaint)
	} else {
		return ansi.Color(cleanIcon, cleanPaint)
	}
}

// GetPaint returns the calculated end paint string for the car.
func (c *Car) GetPaint() string {
	if c.paint = os.Getenv("BULLETTRAIN_CAR_GIT_PAINT"); c.paint == "" {
		c.paint = carPaint
	}

	return c.paint
}

// CanShow decides if this car needs to be displayed.
func (c *Car) CanShow() bool {
	cmd := exec.Command("git", "-C", c.Pwd, "rev-parse", "--git-dir")
	cmdOut, _ := cmd.Output()
	if string(cmdOut) != "" {
		return true
	}

	return false
}

func currentHeadName(pwd string) string {
	cmd := exec.Command("git", "-C", pwd, "symbolic-ref", "HEAD")
	ref, err := cmd.Output()
	if err != nil {
		cmd := exec.Command("git", "-C", pwd, "describe", "--tags", "--exact-match", "HEAD")
		ref, err = cmd.Output()
		if err != nil {
			cmd := exec.Command("git", "-C", pwd, "rev-parse", "--short", "HEAD")
			ref, _ = cmd.Output()
		}
	}

	ref = []byte(strings.Replace(string(ref), "refs/heads/", "", 1))

	if len(ref) == 0 {
		return ""
	}

	return strings.TrimRight(string(ref), "\n")
}

// Render builds and passes the end product of a completely composed car onto
// the channel.
func (c *Car) Render(out chan<- string) {
	defer close(out) // Always close the channel!
	carPaint := ansi.ColorFunc(c.GetPaint())

	out <- fmt.Sprintf("%s%s%s",
		paintedSymbol(),
		carPaint(currentHeadName(c.Pwd)),
		paintStatus(c.Pwd))
}

// GetSeparatorPaint overrides the Fg/Bg colours of the right hand side
// separator through ENV variables.
func (c *Car) GetSeparatorPaint() string {
	return os.Getenv("BULLETTRAIN_CAR_GIT_SEPARATOR_PAINT")
}

// GetSeparatorSymbol overrides the symbol of the right hand side
// separator through ENV variables.
func (c *Car) GetSeparatorSymbol() string {
	return os.Getenv("BULLETTRAIN_CAR_GIT_SEPARATOR_SYMBOL")
}

// GetSeparatorTemplate overrides the template of the right hand side
// separator through ENV variable.
func (c *Car) GetSeparatorTemplate() string {
	return os.Getenv("BULLETTRAIN_CAR_GIT_SEPARATOR_TEMPLATE")
}
