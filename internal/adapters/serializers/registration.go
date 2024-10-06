package serializers

import (
	"encoding/json"

	"github.com/NikolayStrekalov/practicum-gophermart/internal/models"
)

func NewRegistrationFromJSON(data []byte) (*models.Registration, error) {
	r := &models.Registration{}
	err := json.Unmarshal(data, &r)
	if err != nil {
		return nil, err
	}
	return r, nil
}
