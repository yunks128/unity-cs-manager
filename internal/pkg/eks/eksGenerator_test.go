package eks

import (
	"testing"
)

// Explanation of the test:

// - The `tests` variable defines one test case with valid input.
// - The `for` loop iterates over the test cases and runs each test case as a sub-test.
// - For each test case, the `Generate` function is called with the input, and the error returned is checked for being `nil`.
// - If the error returned by `Generate` is not `nil`, the test fails with an error message.
func TestGenerate(t *testing.T) {
	tests := []struct {
		name      string
		nameInput string
		typeInput string
		groups    []NodeGroup
		tags      AWSTags
	}{
		{
			name:      "Valid input",
			nameInput: "test-cluster",
			typeInput: "t2.medium",
			groups: []NodeGroup{
				{
					NodeGroupName:          "worker",
					ClusterInstanceType:    "t2.micro",
					ClusterMinSize:         1,
					ClusterMaxSize:         2,
					ClusterDesiredCapacity: 1,
				},
			},
			tags: AWSTags{
				Servicename:  "dev",
				Creatoremail: "me",
				Resourcename: "unknown",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Generate(tt.nameInput, tt.typeInput, "me", tt.groups, tt.tags)
			if err != nil {
				t.Errorf("Generate() error = %v", err)
			}
		})
	}
}
