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

### Docker compose

Testing 
docker compose test

Running the admin API
docker compose up

### Skaffold (k8s like)

skaffold dev ska-api-admin.yaml

skaffold test ska-api-admin.yaml --profile test-integration

### Deploy

terraform deploy
