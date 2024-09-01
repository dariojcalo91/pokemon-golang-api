package repository

import "context"

func HandleError(ctx context.Context, msg string, err error) error {
	// TODO: should use a logging sistem, maybe Logrus?
	// to handle error specific messaging and data from ctx
	return err
}
