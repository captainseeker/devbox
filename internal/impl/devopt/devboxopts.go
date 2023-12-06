package devopt

import (
	"io"
)

type Opts struct {
	Dir                      string
	Env                      map[string]string
	PreservePathStack        bool
	Pure                     bool
	IgnoreWarnings           bool
	CustomProcessComposeFile string
	Stderr                   io.Writer
}

type GenerateOpts struct {
	Force    bool
	RootUser bool
}

type EnvFlags struct {
	EnvMap  map[string]string
	EnvFile string
}

type PullboxOpts struct {
	Overwrite   bool
	URL         string
	Credentials Credentials
}

type Credentials struct {
	IDToken string
	// TODO We can just parse these out, but don't want to add a dependency right now
	Email string
	Sub   string
}

type AddOpts struct {
	AllowInsecure    bool
	Platforms        []string
	ExcludePlatforms []string
	DisablePlugin    bool
	PatchGlibc       bool
}

type UpdateOpts struct {
	Pkgs                  []string
	IgnoreMissingPackages bool
}

type NixEnvOpts struct {
	DontRecomputeEnvironment bool
	NoRefreshAlias           bool
	RunHooks                 bool
}
