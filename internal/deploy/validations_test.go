package deploy

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestErrors(t *testing.T) {
	require.True(t, isMissing(newMissingError("oops")))
	require.False(t, isValid(newInvalidValueError("oops")))
	require.False(t, isValid(newInvalidUsageError("oops")))
	require.True(t, isMissing(fmt.Errorf("some error %w", newMissingError("oops"))))
	require.False(t, isValid(fmt.Errorf("some error %w", newInvalidValueError("oops"))))
	require.False(t, isValid(fmt.Errorf("some error %w", newInvalidUsageError("oops"))))
}
