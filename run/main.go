package main

import (
	sourceremoval "github.com/initializ-buildpacks/source-removal"
	"github.com/paketo-buildpacks/packit/v2"
)

func main() {
	packit.Run(
		sourceremoval.Detect(),
		sourceremoval.Build())
}
