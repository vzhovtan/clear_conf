package main

import "testing"

func TestCheckToRemove(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"** Policy name: noi-loopback-outbound", false},
		{"*/", false},
		{" description us-nyc-noifw1;", false},
		{"use-interface-description", true},
		{"10.138.0.14;", true},
	}
	for _, test := range tests {
		if got := checkToRemove(test.input); got != test.want {
			t.Errorf("checkToRemove(%q) = %v", test.input, got)
		}
	}
}

func TestValidIPv4(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"10.10.10.1", true},
		{"192.168.1.1/32", true},
		{"172.16.1.1/24", true},
		{"0.256.345.1", false},
	}
	for _, test := range tests {
		if got := validIP4add(test.input); got != test.want {
			t.Errorf("validIP4add(%q) = %v", test.input, got)
		}
	}
}

func TestValidIPv6(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"2a00:79e1:f00::", true},
		{"2a00:79e1:f03:1::/64", true},
		{"2a00:79e1:f03:981::1", true},
		{"2a00:79e1:::f03:981::1", false},
	}
	for _, test := range tests {
		if got := validIP6add(test.input); got != test.want {
			t.Errorf("validIP6add(%q) = %v", test.input, got)
		}
	}
}

func TestWrapOver(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{"10", "10"},
		{"244", "244"},
		{"300", "44"},
		{"512", "0"},
	}
	for _, test := range tests {
		if got := wrapOver(test.input); got != test.want {
			t.Errorf("wrapOver(%q) = %v", test.input, got)
		}
	}
}
