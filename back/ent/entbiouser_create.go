// Code generated by entc, DO NOT EDIT.

package ent

import (
	"back/ent/entbiouser"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// EntBioUserCreate is the builder for creating a EntBioUser entity.
type EntBioUserCreate struct {
	config
	mutation *EntBioUserMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (ebuc *EntBioUserCreate) SetName(s string) *EntBioUserCreate {
	ebuc.mutation.SetName(s)
	return ebuc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ebuc *EntBioUserCreate) SetNillableName(s *string) *EntBioUserCreate {
	if s != nil {
		ebuc.SetName(*s)
	}
	return ebuc
}

// Mutation returns the EntBioUserMutation object of the builder.
func (ebuc *EntBioUserCreate) Mutation() *EntBioUserMutation {
	return ebuc.mutation
}

// Save creates the EntBioUser in the database.
func (ebuc *EntBioUserCreate) Save(ctx context.Context) (*EntBioUser, error) {
	var (
		err  error
		node *EntBioUser
	)
	ebuc.defaults()
	if len(ebuc.hooks) == 0 {
		if err = ebuc.check(); err != nil {
			return nil, err
		}
		node, err = ebuc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EntBioUserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ebuc.check(); err != nil {
				return nil, err
			}
			ebuc.mutation = mutation
			if node, err = ebuc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ebuc.hooks) - 1; i >= 0; i-- {
			if ebuc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ebuc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ebuc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ebuc *EntBioUserCreate) SaveX(ctx context.Context) *EntBioUser {
	v, err := ebuc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ebuc *EntBioUserCreate) Exec(ctx context.Context) error {
	_, err := ebuc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ebuc *EntBioUserCreate) ExecX(ctx context.Context) {
	if err := ebuc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ebuc *EntBioUserCreate) defaults() {
	if _, ok := ebuc.mutation.Name(); !ok {
		v := entbiouser.DefaultName
		ebuc.mutation.SetName(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ebuc *EntBioUserCreate) check() error {
	if _, ok := ebuc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "name"`)}
	}
	return nil
}

func (ebuc *EntBioUserCreate) sqlSave(ctx context.Context) (*EntBioUser, error) {
	_node, _spec := ebuc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ebuc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (ebuc *EntBioUserCreate) createSpec() (*EntBioUser, *sqlgraph.CreateSpec) {
	var (
		_node = &EntBioUser{config: ebuc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: entbiouser.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: entbiouser.FieldID,
			},
		}
	)
	if value, ok := ebuc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: entbiouser.FieldName,
		})
		_node.Name = value
	}
	return _node, _spec
}

// EntBioUserCreateBulk is the builder for creating many EntBioUser entities in bulk.
type EntBioUserCreateBulk struct {
	config
	builders []*EntBioUserCreate
}

// Save creates the EntBioUser entities in the database.
func (ebucb *EntBioUserCreateBulk) Save(ctx context.Context) ([]*EntBioUser, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ebucb.builders))
	nodes := make([]*EntBioUser, len(ebucb.builders))
	mutators := make([]Mutator, len(ebucb.builders))
	for i := range ebucb.builders {
		func(i int, root context.Context) {
			builder := ebucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EntBioUserMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ebucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ebucb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ebucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ebucb *EntBioUserCreateBulk) SaveX(ctx context.Context) []*EntBioUser {
	v, err := ebucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ebucb *EntBioUserCreateBulk) Exec(ctx context.Context) error {
	_, err := ebucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ebucb *EntBioUserCreateBulk) ExecX(ctx context.Context) {
	if err := ebucb.Exec(ctx); err != nil {
		panic(err)
	}
}