package main

import (
	"github.com/unity-sds/unity-cs-terraform-transformer/internal/pkg/eks"
	"reflect"
	"testing"
)

// Explanation of the test:
// - The `tests` variable defines three test cases with different inputs and expected outputs.
// 	- The first test case represents valid input and expects no error and a specific output.
// 	- The second test case represents input with a missing field and expects an error and a `nil` output.
// 	- The third test case represents input with an invalid number and expects an error and a `nil` output.
// 	- The `for` loop iterates over the test cases and runs each test case as a sub-test.
// 	- For each test case, the `arrayToNodeGroup` function is called with the input, and the output is compared with the expected output using `reflect.DeepEqual`.
// 	- If the error returned by `arrayToNodeGroup` doesn't match the expected error, the test fails with an error message.
// - If the output returned by `arrayToNodeGroup` doesn't match the expected output, the test fails with an error message.

func TestArrayToNodeGroup(t *testing.T) {
	tests := []struct {
		name    string
		input   []string
		want    []eks.NodeGroup
		wantErr bool
	}{
		{
			name: "Valid input",
			input: []string{
				"ng1,1,3,2,t2.micro",
				"ng2,2,5,3,t3.medium",
			},
			want: []eks.NodeGroup{
				{
					NodeGroupName:          "ng1",
					ClusterMinSize:         1,
					ClusterMaxSize:         3,
					ClusterDesiredCapacity: 2,
					ClusterInstanceType:    "t2.micro",
				},
				{
					NodeGroupName:          "ng2",
					ClusterMinSize:         2,
					ClusterMaxSize:         5,
					ClusterDesiredCapacity: 3,
					ClusterInstanceType:    "t3.medium",
				},
			},
			wantErr: false,
		},
		{
			name: "Invalid input (missing field)",
			input: []string{
				"ng1,1,3,2",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Invalid input (invalid number)",
			input: []string{
				"ng1,1,3,foo,t2.micro",
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := arrayToNodeGroup(tt.input)

			if (err != nil) != tt.wantErr {
				t.Errorf("arrayToNodeGroup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("arrayToNodeGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}
