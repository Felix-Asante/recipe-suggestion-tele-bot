package ai

import (
	"context"
	"errors"

	"github.com/replicate/replicate-go"
)

func GeneratePhotos(text string) (string, error) {
	ctx := context.TODO()
	r8, err := replicate.NewClient(replicate.WithTokenFromEnv())
	if err != nil {
		return "", err
	}

	model := "black-forest-labs/flux-schnell"

	input := replicate.PredictionInput{
		"prompt": text,
	}

	// Run a model and wait for its output
	output, err := r8.Run(ctx, model, input, nil)
	if err != nil {
		return "", err
	}

	url, ok := output.([]interface{})

	if !ok {
		return "", errors.New("invalid output")
	}

	return url[0].(string), nil
}
