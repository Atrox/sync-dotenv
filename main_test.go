package main

import (
	"os"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

const dotEnvFile = `
# COMMENT 1
KEY1=VALUE1
KEY2="VALUE2"

# COMMENT 2
KEY3=VALUE3

# COMMENT 3
KEY4=VALUE4
`

const dotEnvExampleFile = `
# COMMENT 3
KEY4=

# COMMENT 2
KEY3=EXAMPLEVALUE
`

const dotEnvExpected = `
# COMMENT 1
KEY1=
KEY2=

# COMMENT 2
KEY3=EXAMPLEVALUE

# COMMENT 3
KEY4=
`

func init() {
	envFilePath = ".env"
	exampleFilePath = ".env.example"
}

func TestRun(t *testing.T) {
	fs = afero.NewMemMapFs()

	err := afero.WriteFile(fs, envFilePath, []byte(dotEnvFile), os.ModePerm)
	assert.Nil(t, err)

	err = afero.WriteFile(fs, exampleFilePath, []byte(dotEnvExampleFile), os.ModePerm)
	assert.Nil(t, err)

	err = run(nil, nil)
	assert.Nil(t, err)

	fileContent, err := afero.ReadFile(fs, exampleFilePath)
	assert.Nil(t, err)

	assert.Equal(t, dotEnvExpected, string(fileContent), "they should be equal")
}

func TestWithoutEnvFile(t *testing.T) {
	fs = afero.NewMemMapFs()

	err := run(nil, nil)
	assert.NotNil(t, err)
}
