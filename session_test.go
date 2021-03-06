package gitbase

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/connectivity"
)

func TestSessionBblfshClient(t *testing.T) {
	require := require.New(t)

	session := NewSession(nil, WithBblfshEndpoint(defaultBblfshEndpoint))
	cli, err := session.BblfshClient()
	require.NoError(err)
	require.NotNil(cli)
	require.Equal(connectivity.Ready, cli.GetState())
}

func TestSupportedLanguagesAliases(t *testing.T) {
	require := require.New(t)

	session := NewSession(nil, WithBblfshEndpoint(defaultBblfshEndpoint))
	cli, err := session.BblfshClient()
	require.NoError(err)
	require.NotNil(cli)
	require.Equal(connectivity.Ready, cli.GetState())
	ok, err := cli.IsLanguageSupported(context.TODO(), "C++")
	require.NoError(err)
	require.True(ok)
}

func TestSessionBblfshClientNoConnection(t *testing.T) {
	require := require.New(t)

	session := NewSession(nil, WithBblfshEndpoint("localhost:9999"))
	_, err := session.BblfshClient()
	require.Error(err)
	require.True(ErrBblfshConnection.Is(err))
}
