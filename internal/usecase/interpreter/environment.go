package interpreter

import (
	"fmt"

	"github.com/avazquezcode/govetryx/internal/domain/types"
)

type Env struct {
	values types.HashMap
	parent *Env
}

// NewLocal is a constructor for a global environment (env without parent).
func NewGlobal() *Env {
	return &Env{
		values: types.HashMap{},
	}
}

// NewLocal is a constructor for a local environment (env that has a parent).
func NewLocal(parent *Env) *Env {
	return &Env{
		values: types.HashMap{},
		parent: parent,
	}
}

// Get is a getter to get the value of a key within the environment.
// It tries to find the key first on the local environment,
// and recursivelly do the same in all the parent envs until finding it.
func (e *Env) Get(key string) (interface{}, error) {
	if e.values.Exists(key) {
		return e.values.Get(key), nil
	}

	// try recursively on the parents
	if e.parent != nil {
		return e.parent.Get(key)
	}

	return nil, fmt.Errorf("the variable %s is not defined", key)
}

// Set sets a new entry in the environment (key -> value)
func (e *Env) Set(key string, value interface{}) {
	e.values.Set(key, value)
}

// Assigns a value to an "already declared" variable.
// It tries to find the key first on the local environment,
// and recursivelly do the same in all the parent envs until finding it,
// to proceed with the assignment if found.
func (e *Env) Assign(key string, value interface{}) error {
	if e.values.Exists(key) {
		e.values.Set(key, value)
		return nil
	}

	// check recursively for the parents
	if e.parent != nil {
		return e.parent.Assign(key, value)
	}

	return fmt.Errorf("cannot assign a value to the variable %q, because it is not declared", key)
}

// GetAt try to get the key value in a specific depth.
func (e *Env) GetAt(depth int, key string) (interface{}, error) {
	ancestor := e.ancestor(depth)
	if !ancestor.values.Exists(key) {
		return nil, fmt.Errorf("could not find the variable %s", key)
	}

	return ancestor.values.Get(key), nil
}

func (e *Env) ancestor(depth int) *Env {
	env := e
	for i := 0; i < depth; i++ {
		env = env.parent
	}
	return env
}

// AssignAt try to assign the key value in a specific depth.
func (e *Env) AssignAt(depth int, key string, value interface{}) error {
	e.ancestor(depth).values.Set(key, value)
	return nil
}
