package user

import (
	"problem2/user/course"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitializeUserData(t *testing.T) {

	tests := []struct {
		users           []User
		expectedUserMap map[int]User
	}{
		{
			[]User{
				{"A", 1, "Addr1", 10, []course.Course{"A", "B", "C", "D"}},
				{"B", 2, "Addr2", 11, []course.Course{"A", "B", "C", "E"}},
			},
			map[int]User{
				10: {"A", 1, "Addr1", 10, []course.Course{"A", "B", "C", "D"}},
				11: {"B", 2, "Addr2", 11, []course.Course{"A", "B", "C", "E"}},
			},
		},
	}

	for _, tc := range tests {
		InitializeUserData(tc.users)
		assert.Equal(t, tc.expectedUserMap, UserMap, "validate InitializeUserData")
	}
}
