// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package di

import (
	"github.com/google/wire"
	"github.com/int128/kubelogin/adaptors"
	"github.com/int128/kubelogin/adaptors/cmd"
	"github.com/int128/kubelogin/adaptors/kubeconfig"
	"github.com/int128/kubelogin/adaptors/logger"
	"github.com/int128/kubelogin/adaptors/oidc"
	"github.com/int128/kubelogin/usecases"
	"github.com/int128/kubelogin/usecases/login"
)

// Injectors from di.go:

func NewCmd() adaptors.Cmd {
	kubeconfigKubeconfig := &kubeconfig.Kubeconfig{}
	adaptorsLogger := logger.New()
	factory := &oidc.Factory{
		Logger: adaptorsLogger,
	}
	prompt := &login.Prompt{
		Logger: adaptorsLogger,
	}
	loginLogin := &login.Login{
		Kubeconfig: kubeconfigKubeconfig,
		OIDC:       factory,
		Logger:     adaptorsLogger,
		Prompt:     prompt,
	}
	cmdCmd := &cmd.Cmd{
		Login:  loginLogin,
		Logger: adaptorsLogger,
	}
	return cmdCmd
}

func NewCmdWith(adaptorsLogger adaptors.Logger, loginPrompt usecases.LoginPrompt) adaptors.Cmd {
	kubeconfigKubeconfig := &kubeconfig.Kubeconfig{}
	factory := &oidc.Factory{
		Logger: adaptorsLogger,
	}
	loginLogin := &login.Login{
		Kubeconfig: kubeconfigKubeconfig,
		OIDC:       factory,
		Logger:     adaptorsLogger,
		Prompt:     loginPrompt,
	}
	cmdCmd := &cmd.Cmd{
		Login:  loginLogin,
		Logger: adaptorsLogger,
	}
	return cmdCmd
}

// di.go:

var usecasesSet = wire.NewSet(login.Login{}, wire.Bind((*usecases.Login)(nil), (*login.Login)(nil)))

var adaptorsSet = wire.NewSet(cmd.Cmd{}, kubeconfig.Kubeconfig{}, oidc.Factory{}, wire.Bind((*adaptors.Cmd)(nil), (*cmd.Cmd)(nil)), wire.Bind((*adaptors.Kubeconfig)(nil), (*kubeconfig.Kubeconfig)(nil)), wire.Bind((*adaptors.OIDC)(nil), (*oidc.Factory)(nil)))

var extraSet = wire.NewSet(login.Prompt{}, wire.Bind((*usecases.LoginPrompt)(nil), (*login.Prompt)(nil)), logger.New)