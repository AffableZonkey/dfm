package dfm

import "github.com/spf13/afero"

// CONFIG holds the global config struct
var CONFIG Config

// DRYRUN is used to globally set whether this is a dry run
var DRYRUN = false

var FS = afero.NewOsFs()
