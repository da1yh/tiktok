package dao

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func SLEEP() {
	time.Sleep(3 * time.Second)
}

// 4 -> 14 'hello s1mple, i am zywoo'
// 4 -> 16 'hello niko, i am zywoo'
// 14 -> 4 'hi zywoo, i am s1mple'
// 4 -> 14 'you are 🐐'
func TestMessageDao(t *testing.T) {
	InitDb()
	id1, err := AddMessageByAll(4, 14, "hello s1mple, i am zywoo", time.Now())
	assert.Nil(t, err)

	SLEEP()

	id2, err := AddMessageByAll(4, 16, "hello niko, i am zywoo", time.Now())
	assert.Nil(t, err)

	SLEEP()

	id3, err := AddMessageByAll(14, 4, "hi zywoo, i am s1mple", time.Now())
	assert.Nil(t, err)

	SLEEP()

	id4, err := AddMessageByAll(4, 14, "you are 🐐", time.Now())

	ids, err := FindMessageIdsByFromUserIdAndToUserId(4, 14)
	assert.Nil(t, err)
	assert.Equal(t, len(ids), 2)
	assert.True(t, (ids[0] == id1 && ids[1] == id4) || (ids[0] == id4 && ids[1] == id1))

	message, err := FindMessageById(id2)
	assert.Nil(t, err)
	assert.Equal(t, message.Id, id2)
	assert.Equal(t, message.FromUserId, int64(4))
	assert.Equal(t, message.ToUserId, int64(16))
	assert.Equal(t, message.Content, "hello niko, i am zywoo")

	message, err = FindMessageById(id3)
	assert.Nil(t, err)
	assert.Equal(t, message.Id, id3)
	assert.Equal(t, message.FromUserId, int64(14))
	assert.Equal(t, message.ToUserId, int64(4))
	assert.Equal(t, message.Content, "hi zywoo, i am s1mple")

}