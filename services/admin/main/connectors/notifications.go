package connectors

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
)

func NotifyExcessOfDevices(abbr string, message string) (*http.Response, error) {
	println("Notifying")
	var host = os.Getenv("NOTIFICATIONS_SERVICE_HOST")
	var port = os.Getenv("NOTIFICATIONS_SERVICE_PORT")
	body := map[string]string{"level": "warning", "employeeAbbreviation": abbr, "message": message}
	jsonValue, _ := json.Marshal(body)
	request, err := http.NewRequest("POST", "http://"+host+":"+port+"/api/notify", bytes.NewBuffer(jsonValue))
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		panic(error)
	}
	defer response.Body.Close()
	if err != nil {
		// TODO: Bugsnag or sentry this error case
		println(err.Error())
		return nil, err
	}
	return response, err
}
