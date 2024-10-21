package biz

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

// Task entity.
type Task struct {
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at"`
	// NillableCreatedAt holds the value of the "nillable_created_at" field.
	NillableCreatedAt *time.Time `json:"nillable_created_at,omitempty"`
}

func TestStat(t *testing.T) {
	var name = "abc"
	var age int
	assert.NotZero(t, name)
	assert.Zero(t, age)

	b, _ := json.Marshal(Task{})
	fmt.Printf("%s\n", b)
	// {"created_at":"0001-01-01T00:00:00Z"}

	now := time.Now()
	b, _ = json.Marshal(Task{CreatedAt: now, NillableCreatedAt: &now})
	fmt.Printf("%s\n", b)
	// {"created_at":"2009-11-10T23:00:00Z","nillable_created_at":"2009-11-10T23:00:00Z"}
}
