package main

import (
	"os"
	"testing"
	"time"

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

const dotEnvFileWithChanges = `
# COMMENT 2
KEY3=VALUE3

# COMMENT 1
KEY1=VALUE1
KEY2="VALUE2"
`

const dotEnvExpectedWithChanges = `
# COMMENT 2
KEY3=EXAMPLEVALUE

# COMMENT 1
KEY1=
KEY2=
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

func TestWatch(t *testing.T) {
	fs = afero.NewOsFs()
	watch = true

	err := afero.WriteFile(fs, envFilePath, []byte(dotEnvFile), os.ModePerm)
	assert.Nil(t, err)

	err = afero.WriteFile(fs, exampleFilePath, []byte(dotEnvExampleFile), os.ModePerm)
	assert.Nil(t, err)

	go func() {
		err = run(nil, nil)
	}()

	time.Sleep(100 * time.Millisecond)

	err = afero.WriteFile(fs, envFilePath, []byte(dotEnvFileWithChanges), os.ModePerm)
	assert.Nil(t, err)

	time.Sleep(100 * time.Millisecond)

	fileContent, err := afero.ReadFile(fs, exampleFilePath)
	assert.Nil(t, err)

	assert.Equal(t, dotEnvExpectedWithChanges, string(fileContent), "they should be equal")
}

func TestWithoutEnvFile(t *testing.T) {
	fs = afero.NewMemMapFs()

	err := run(nil, nil)
	assert.NotNil(t, err)
}
