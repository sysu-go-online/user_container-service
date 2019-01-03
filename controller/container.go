package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sysu-go-online/public-service/tools"
	"github.com/sysu-go-online/user_container-service/model"
)

// ChangePortStatusHandler open or close a port of given container id
func ChangePortStatusHandler(w http.ResponseWriter, r *http.Request) error {

}

// GetPortsStatusHandler get ports status of given container id
func GetPortsStatusHandler(w http.ResponseWriter, r *http.Request) error {

}

// ChangeContainerStatusHandler turn on or turn off a container with given id
func ChangeContainerStatusHandler(w http.ResponseWriter, r *http.Request) error {

}

// CreateContainerHandler create an empty with given message for user
func CreateContainerHandler(w http.ResponseWriter, r *http.Request) error {
	// Check token
	if ok, err := tools.CheckJWT(r.Header.Get("Authorization"), AuthRedisClient); !(ok && err == nil) {
		w.WriteHeader(401)
		return nil
	}
	r.ParseForm()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	// TODO: parse jwt to get username
	username := mux.Vars(r)["username"]

	ret := ClientCreateContainerResponse{}
	if model.CheckRemainingContainerAmount(username) <= 0 {
		ret.OK = false
		ret.Msg = "exert the max number"
		byteRet, err := json.Marshal(&ret)
		if err != nil {
			return err
		}
		w.Write(byteRet)
		return nil
	}

	jsonBody := ClientCreateContainerRequest{}
	err = json.Unmarshal(body, &jsonBody)
	if err != nil {
		w.WriteHeader(400)
		return nil
	}
	id, err := createContainer(jsonBody.BaseImage)
	if err != nil {
		return err
	}

	// TODO: use transaction
	err = model.MinusRemainedContainerAmount(username)
	if err != nil {
		// TODO: destory created container
		return err
	}
	err = model.RecordContainerForUser(username, id)
	if err != nil {
		// TODO: destory created container
		return err
	}

	return nil
}

// RemoveContainerHandler delete a container with given id
func RemoveContainerHandler(w http.ResponseWriter, r *http.Request) error {

}

// GetAllContainersStatusHandler give open status of all of the container
func GetAllContainersStatusHandler(w http.ResponseWriter, r *http.Request) error {

}

// TODO:
func createContainer(imageName string) (string, error) {

}
