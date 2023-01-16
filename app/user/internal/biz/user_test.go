package biz

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/stretchr/testify/assert"
	"kratos-realworld/pkg/middleware/auth"
	"testing"
)

func Test_HashPassword(t *testing.T) {
	// password=abc , hashedPwd=$2a$10$K8nTiZD8H5KSnf/QIf2TCuMmfX21man00Fd56ZT.mqfAzEKY//FEy
	// password=abc , hashedPwd=$2a$10$z.VPneiaUo4LD.EUanASme1MpE8tCC/FFdNL2WnCP9ECO3fSLmkXa
	hashPassword("abc")
}

func Test_VerifyPassword(t *testing.T) {
	testCases := []struct {
		name      string
		input     string
		hashedPwd string
		result    bool
	}{
		{
			name:      "verity success",
			input:     "abc",
			hashedPwd: "$2a$10$K8nTiZD8H5KSnf/QIf2TCuMmfX21man00Fd56ZT.mqfAzEKY//FEy",
			result:    true,
		},
		{
			name:      "verity success",
			input:     "abc",
			hashedPwd: "$2a$10$z.VPneiaUo4LD.EUanASme1MpE8tCC/FFdNL2WnCP9ECO3fSLmkXa",
			result:    true,
		},
		{
			name:      "input not right",
			input:     "abc2",
			hashedPwd: "$2a$10$z.VPneiaUo4LD.EUanASme1MpE8tCC/FFdNL2WnCP9ECO3fSLmkXa",
			result:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := verifyPassword(tc.hashedPwd, tc.input)
			assert.Equal(t, tc.result, result)
		})
	}
}

func Test_GenerateToken(t *testing.T) {
	user := &auth.User{
		Id:       1,
		Username: "zhang3",
		Email:    "zhang3@126.com",
	}
	tokenString := auth.GenerateToken("secret", user)
	log.Info(tokenString)
}
