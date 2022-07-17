package balances

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHasOption(t *testing.T) {

	require.True(t, (VIRTUAL | PASSIVE).IsPassive())
	require.True(t, (VIRTUAL | PASSIVE).IsVirtual())

	require.False(t, (VIRTUAL | PASSIVE).IsSyncable())
	require.False(t, (VIRTUAL | PASSIVE).IsActive())
}
