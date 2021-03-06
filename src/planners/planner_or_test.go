// Copyright 2020 The VectorSQL Authors.
//
// Code is licensed under Apache License, Version 2.0.

package planners

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrPlan(t *testing.T) {
	plan := NewOrPlan(
		NewBooleanExpressionPlan(
			"=",
			NewVariablePlan("name"),
			NewConstantPlan("x"),
		),
		NewBooleanExpressionPlan(
			"=",
			NewVariablePlan("name"),
			NewConstantPlan("y"),
		),
	)
	err := plan.Build()
	assert.Nil(t, err)
	t.Logf("%v", plan.Name())

	err = plan.Walk(func(plan IPlan) (bool, error) {
		return true, nil
	})
	assert.Nil(t, err)

	expect := "OrNode=(Func=[OR], Left=[BooleanExpressionNode=(Func=[=], Args=[[VariableNode=[$name] ConstantNode=<x>]])], Right=[BooleanExpressionNode=(Func=[=], Args=[[VariableNode=[$name] ConstantNode=<y>]])])"
	actual := plan.String()
	assert.Equal(t, expect, actual)
}
