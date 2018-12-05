package main

import "testing"
import "github.com/stretchr/testify/assert"

func TestConvertToInt(t *testing.T) {
	assert.Equal(t, 1, convertToInt('a'))
	assert.Equal(t, -1, convertToInt('A'))
	assert.Equal(t, 19, convertToInt('s'))
	assert.Equal(t, -19, convertToInt('S'))
	assert.Equal(t, 26, convertToInt('z'))
	assert.Equal(t, -26, convertToInt('Z'))
}

func TestConvertToIntArray(t *testing.T) {
	assert.Equal(t, []int{1, 19, 26, -1, -19, -26}, convertToIntArray([]byte("aszASZ")))
}

func TestDestroyFirstPair(t *testing.T) {
	assert.Equal(t, []int{1, 19, 26, -1, -19, -26}, destroyFirstPair([]int{1, 19, 26, -1, -19, -26}))
	assert.Equal(t, []int{}, destroyFirstPair([]int{}))
	assert.Equal(t, []int{2}, destroyFirstPair([]int{2}))
	assert.Equal(t, []int{}, destroyFirstPair([]int{1, -1}))
	assert.Equal(t, []int{1, 1}, destroyFirstPair([]int{1, 1, -1, 1}))
	assert.Equal(t, []int{1, -1, 3}, destroyFirstPair([]int{1, -2, 2, -1, 3}))
}

func TestStripUnitType(t *testing.T) {
	assert.Equal(t, []int{1, 26, -1, -26}, stripUnitType([]int{1, 19, 26, -1, -19, -26}, 19))
	assert.Equal(t, []int{}, stripUnitType([]int{}, 1))
	assert.Equal(t, []int{2}, stripUnitType([]int{2}, 1))
	assert.Equal(t, []int{3}, stripUnitType([]int{1, 1, 3}, 1))
	assert.Equal(t, []int{3}, stripUnitType([]int{1, -1, 3}, 1))
	assert.Equal(t, []int{}, stripUnitType([]int{1, -1}, 1))
	assert.Equal(t, []int{}, stripUnitType([]int{1, 1, -1, 1}, 1))
	assert.Equal(t, []int{1, -1, 3}, stripUnitType([]int{1, -2, 2, -1, 3}, 2))
}
