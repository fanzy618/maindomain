package maindomain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMainDomain(t *testing.T) {
	a := assert.New(t)
	cases := [...][2]string{
		{"*.lb-ns.myalauda.cn", "myalauda.cn"},
		{"sina.com.cn", "sina.com.cn"},
		{"www.alauda.io", "alauda.io"},
		{"www.ele.me", "ele.me"},
		{"t.cn", "t.cn"},
		{"bbs.hust.edu.cn", "hust.edu.cn"},
		{"image.google.com.hk", "google.com.hk"},
	}
	for _, testCase := range cases {
		a.Equal(testCase[1], FindMainDomain(testCase[0]))
	}
}
