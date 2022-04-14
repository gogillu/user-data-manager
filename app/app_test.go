package app

/*
func TestValidateNewUserDetails(t *testing.T) {

	positiveTests := []struct {
		user           user.User
		expectedOutput bool
	}{
		{
			user.User{"A", 1, "Addr1", 10, []course.Course{"A", "B", "C", "D"}},
			true,
		},
	}

	for _, tc := range positiveTests {
		assert.Equal(t, tc.expectedOutput, validateNewUserDetails(tc.user), "validate new User")
	}

	negativeTests := []struct {
		user           user.User
		expectedOutput bool
	}{
		{
			user.User{"A", 1, "Addr1", 10, []course.Course{"A", "B", "C"}},
			true,
		},
		{
			user.User{"A", 1, "Addr1", 10, []course.Course{"A", "B", "C", "D", "E", "F", "A"}},
			true,
		},
	}

	for _, tc := range negativeTests {
		assert.NotEqual(t, tc.expectedOutput, validateNewUserDetails(tc.user), "validate new User")
	}

}
*/
