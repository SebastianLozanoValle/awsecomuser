package secretm

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/sebastianlozanovalle/awsecomuser/awsgo"
	"github.com/sebastianlozanovalle/awsecomuser/models"
)

func GetSecret(secretName string) (models.SecretRDSJson, error) {
	var dataSecret models.SecretRDSJson
	fmt.Println(" > Pido secreto " + secretName)

	svc := secretsmanager.NewFromConfig(awsgo.Cfg)
	clave, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	})

	if err != nil {
		fmt.Println(err.Error())
		return dataSecret, err
	}

	json.Unmarshal([]byte(*clave.SecretString), &dataSecret)
	fmt.Println(" > Lectura Secret OK " + secretName)
	return dataSecret, nil
}