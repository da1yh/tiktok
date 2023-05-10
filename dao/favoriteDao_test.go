package dao

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInit3(t *testing.T) {
	InitDb()
}

func TestFavorite(t *testing.T) {
	err := AddFavoriteByBothId(4, 2)
	assert.Nil(t, err)
	err = AddFavoriteByBothId(14, 2)
	assert.Nil(t, err)
	err = AddFavoriteByBothId(16, 2)
	assert.Nil(t, err)
	num, err := CountFavoritesByToVideoId(2)
	assert.Nil(t, err)
	assert.Equal(t, num, int64(3))
	res, err := CheckFavoriteByBothId(4, 2)
	assert.Nil(t, err)
	assert.True(t, res)
	err = DeleteFavoriteByBothId(4, 2)
	assert.Nil(t, err)
	res, err = CheckFavoriteByBothId(4, 2)
	assert.Nil(t, err)
	assert.False(t, res)
}