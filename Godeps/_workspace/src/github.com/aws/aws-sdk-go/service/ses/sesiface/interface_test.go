// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package sesiface_test

import (
	"testing"

	"github.com/rjocoleman/lambda-ses/Godeps/_workspace/src/github.com/aws/aws-sdk-go/service/ses"
	"github.com/rjocoleman/lambda-ses/Godeps/_workspace/src/github.com/aws/aws-sdk-go/service/ses/sesiface"
	"github.com/stretchr/testify/assert"
)

func TestInterface(t *testing.T) {
	assert.Implements(t, (*sesiface.SESAPI)(nil), ses.New(nil))
}
