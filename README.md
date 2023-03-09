# Greenbone excercise

This is the system for device control for the different users of the company SampleCompany.

## Features

- Stores the different devices
  - MAC Address
  - Computer Name
  - IP Address
  - Employee Abbreviation
  - Description

- The system administrator wants to be able to add a new computer to an employee
- The system administrator wants to be informed when an employee is assigned 3 or
more computers
- The system administrator wants to be able to get all computers
- The system administrator wants to be able to get all assigned computers for an
employee 
- The system administrator wants to be able to get the data of a single computer
- The system administrator wants to be able to remove a computer from an employee
- The system administrator wants to be able to assign a computer to another employee

## Ways of running the given project

### Prerequisites

- Installation of minikube / kind or k3s.

- Once this installation is done, you may need to run the shell script `install_once.sh`

With this seccomp profile will be added as secure_profile.json

### Skaffold (k8s like)

This was tested out with minikube, but you may choose your own kind, k3s etc, this selection is done inside of the skaffold file (by default minikube but I commented out for your choose)

### Development mode

Runs with hot reload images to develop in the cluster.

$ skaffold dev skaffold_dev.yaml

### Production mode (local)

Runs the production environment in a local cluster (production images)

$ skaffold run skaffold_prod.yaml

### CI mode

Runs the CI mode to test the piece of the cluster you develop locally.

skaffold test skaffold_ci.yaml --profile test-integration

### Environment Variables

Generally speaking they belong to configmaps or inside of secrets.
As a possible improvement secrets could belong to a given volume instead of environment variable, encrypted may be git gitcrypt or using aws secret manager.

### Storage

Currently the saving of information belongs to the time of the "living" pod postgres. In order to store for the time "living" of minikube / kind or other you may want to add a persisten volume of kind host.

Migrations is currently used as a k8s batch job that runs the migrate script to initalize the db schemas.

One possible improvement would be to add FK for user_devices table considering user id and device id.

### Local Test

Once it runs with skaffold dev or prod mode:

Execute this command to see the `minikube service -n services admin --url` to get the exposed node port service admin
Most likely will be a combination of `minikube ip` + the NodePort port 32020 for admin, 32021 for notify (the service you created for alarms)

Seed the base database:

`curl -X POST http://[IP of minikube kind or k3s]:32020/seed`

This will give back similar information to this:

`{"assignedUserOnce":"uuid","assignedUserTwoTimes":"uuid", "deviceAssignedOnce": "uuid", "deviceAssignedTwice": "uuid"}`

so now you can call the following curl / httpie / wget to assign the given device to a user

`curl -X POST http://[IP]:32020/user/[userUUID]/device/[deviceUUID]`

### Documentation

The service admin contains the /docs folder with some documentation in swagger yaml like.

### Tests

One unit testing

One super test

Both run with ci skaffold