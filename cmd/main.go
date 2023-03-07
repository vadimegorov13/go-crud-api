package main

import (
	"fmt"

	"github.com/spf13/viper"
)

/*
Initialize viper to handle env variables
*/

func main() {
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()

	// add env variables as needed
	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)

	fmt.Println(port, dbUrl)

}
