package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct__Failure(t *testing.T) {
	c, err := NewCourse("", "description", "category")
	assert.Nil(t, c)
	assert.NotNil(t, err)
	assert.Error(t, err)
	assert.EqualError(t, err, "title is required")

	c, err = NewCourse("title", "", "category")
	assert.Nil(t, c)
	assert.NotNil(t, err)
	assert.Error(t, err)
	assert.EqualError(t, err, "description is required")

	c, err = NewCourse("title", "description", "")
	assert.Nil(t, c)
	assert.NotNil(t, err)
	assert.Error(t, err)
	assert.EqualError(t, err, "category is required")
}

func TestNewProduct__Success(t *testing.T) {
	c, err := NewCourse("Kafka", "Aprenda Kafka", "Mensageria")
	assert.Nil(t, err)
	assert.NotNil(t, c)
	assert.NotEmpty(t, c.ID)
	assert.Equal(t, c.Title, "Kafka")
	assert.Equal(t, c.Description, "Aprenda Kafka")
	assert.Equal(t, c.Category, "Mensageria")
	assert.NotEmpty(t, c.CreatedAt)
	assert.NotEmpty(t, c.UpdatedAt)
}
