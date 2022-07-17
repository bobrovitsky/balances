package balances

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHasOption(t *testing.T) {

	require.True(t, (VIRTUAL | PASSIVE).isPassive())
	require.True(t, (VIRTUAL | PASSIVE).isVirtual())

	require.False(t, (VIRTUAL | PASSIVE).isSyncable())
	require.False(t, (VIRTUAL | PASSIVE).isActive())
}
