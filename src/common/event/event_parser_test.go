package event

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestEventDataToItem(t *testing.T) {
	t.Run("should parse interface to attributes", func(t *testing.T) {
		type testcase struct {
			ID      string            `json:"id"`
			Name    string            `json:"name"`
			Age     int               `json:"age"`
			Male    bool              `json:"isMale"`
			Today   time.Time         `json:"today"`
			Skills  []string          `json:"skills"`
			Parents map[string]string `json:"parents"`
		}

		fakeID := uuid.NewString()
		fakeCreatedAt := time.Now()
		fakeData := testcase{
			ID:     "123",
			Name:   "John doe",
			Age:    18,
			Male:   true,
			Today:  time.Now(),
			Skills: []string{"Coding like crazy", "Talk shit"},
			Parents: map[string]string{
				"Mother": "Taylor",
				"Father": "Joseph",
			},
		}

		event := NewEvent(
			"ActionPerformed",
			ID(fakeID),
			Data(fakeData),
			CreatedAt(fakeCreatedAt),
		)

		fakeDataBytes, err := json.Marshal(fakeData)

		want := map[string]types.AttributeValue{
			"ID":        &types.AttributeValueMemberS{Value: fakeID},
			"Name":      &types.AttributeValueMemberS{Value: "ActionPerformed"},
			"CreatedAt": &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", fakeCreatedAt.Unix())},
			"Data":      &types.AttributeValueMemberB{Value: fakeDataBytes},
		}

		have := eventDataToItem(event)

		assert.Nil(t, err)
		assert.Equal(t, want["ID"], have["ID"])
		assert.Equal(t, want["Name"], have["Name"])
		assert.Equal(t, want["CreatedAt"], have["CreatedAt"])
		assert.Equal(t, want["Data"], have["Data"])
	})
}
