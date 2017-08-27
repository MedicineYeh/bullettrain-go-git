package carGit

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/bullettrain-sh/bullettrain-go-core/ansi"
)

const (
	carPaint       = "black:white"
	gitSymbolPaint = "red:white"
	gitSymbolIcon  = ""
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

func currentBranchName(pwd string) string {
	cmd := exec.Command(
		"git", "-C", pwd, "rev-parse", "--abbrev-ref", "HEAD")
	cmdOut, _ := cmd.Output()
	if string(cmdOut) == "" {
		return ""
	}

	return strings.TrimRight(string(cmdOut), "\n")
}

// Render builds and passes the end product of a completely composed car onto
// the channel.
func (c *Car) Render(out chan<- string) {
	defer close(out) // Always close the channel!
	carPaint := ansi.ColorFunc(c.GetPaint())

	out <- fmt.Sprintf("%s%s",
		paintedSymbol(),
		carPaint(currentBranchName(c.Pwd)))
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
