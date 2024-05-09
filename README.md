*DRAFT*


---

**Opis**

<p align="center">
  <img src="https://github.com/PiotrFerenc/mash2/assets/30370747/27230488-59c6-48e5-8def-8f345c0ac38e" alt="logo" width="400"/>
</p>
<div style="text-align: justify;">

[![Go](https://github.com/PiotrFerenc/mash2/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/PiotrFerenc/mash2/actions/workflows/go.yml)

[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=PiotrFerenc_mash2&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=PiotrFerenc_mash2)

![100 - commitów](https://img.shields.io/badge/100-commitów-2ea44f?logo=go)

# Task tower

Task Tower is a tool for automating distributed tasks, utilizing pipelines with isolated processing steps in Docker containers, which allows users to easily scale and modify processes without impacting the entire infrastructure.
 
</div>

## Benefits

- Flexibility in task management
- Process isolation
- Increased operational efficiency
- Ease of integration and collaboration

## How it works

```mermaid
sequenceDiagram
    participant a as API
    participant c as Controller
    participant q as Queue
    participant w1 as Worker 1 
    participant w2 as Worker 2
    participant w3 as Worker 3

    a->>c: API Call Request
    c->>a: Response: Pipeline Added to Queue
    c->>q: Add new task: Action1
    q->>w1: Assign Action1 to Worker 1
    w1->>q: Action1 Completed
    q->>c: Notify Controller: Action1 Done
    c->>q: Add new task: Action2
    q->>w2: Assign Action2 to Worker 2
    w2->>q: Action2 Completed
    q->>c: Notify Controller: Action2 Done
    c->>q: Add new task: Action3
    q->>w3: Assign Action3 to Worker 3
    w3->>q: Action3 Completed
    q->>c: Notify Controller: Action3 Done
    c->>q: Add new task: Action4
    q->>w1: Assign Action3 to Worker 1
    w1->>q: Action3 Completed
    q->>c: Notify Controller: Action3 Done
```
The diagram shows the interactions between different participants in a task processing system. The API receives a request from a user and sends a response that the task has been added to the queue by the controller. Then, the controller adds a new task, labeled Action1, to the queue. The queue assigns this task to Worker 1, who, after completing it, informs the queue that it's done. The queue then notifies the controller that Action1 is complete. The process repeats for the next tasks, Action2 and Action3, which are assigned to Worker 2 and Worker 3, respectively. Each of these tasks is completed and reported as finished to the controller by the respective workers. Finally, Action4 is added to the queue, assigned again to Worker 1, who completes the task and reports its completion. The queue again informs the controller that the task is done.

## Installation 
Linux:
```shell
wget https://raw.githubusercontent.com/PiotrFerenc/mash2/main/install.sh && chmod +x install.sh && ./install.sh
```

## Build 

### requirements

- [MAKE](https://www.gnu.org/software/make/)
- [docker](https://docs.docker.com/engine/install/)
- [docker-compose](https://docs.docker.com/compose/install/)
- [GO](https://go.dev/doc/install)

```git
git clone https://github.com/PiotrFerenc/mash2
```

```shell
cd mash2
```

```makefile
make docker-rebuild
```

## Usage 

```shell
curl -X POST localhost:5000/execute -H "Content-Type: application/json" -d '{
    "Parameters": {
        "console.text": "hallo word"
    },
    "Tasks": [
        {
            "Sequence": 1,
            "Name": "log",
            "Action": "console"
        }
    ]
}'
```

or
```shell

echo '{
    "Parameters": {
        "console.text": "hallo word"
    },
    "Tasks": [
        {
            "Sequence": 1,
            "Name": "log",
            "Action": "console"
        }
    ]
}' > request.json

curl -X POST localhost:5000/execute -H "Content-Type: application/json" -d @request.json
```

## Actions


## Todo 1.0
- [x] Application Programming Interface (API)
- [ ] Monitoring and Logging
- [ ] CLI
- [x] Triggering Tasks
- [x] Handling Distributed Tasks
- [x] Error Management
- [x] Extensibility
- [x] Configuration

## Todo: 2.0
- [ ] User Interface (UI)
- [ ] Scheduling
- [ ] Dependency Management of Tasks
- [ ] Retry Mechanisms
- [ ] Security
- [ ] Integration with External Services and Applications

## STATISTICS

[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=PiotrFerenc_mash2&metric=bugs)](https://sonarcloud.io/summary/new_code?id=PiotrFerenc_mash2) [![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=PiotrFerenc_mash2&metric=code_smells)](https://sonarcloud.io/summary/new_code?id=PiotrFerenc_mash2) [![Coverage](https://sonarcloud.io/api/project_badges/measure?project=PiotrFerenc_mash2&metric=coverage)](https://sonarcloud.io/summary/new_code?id=PiotrFerenc_mash2) [![Duplicated Lines (%)](https://sonarcloud.io/api/project_badges/measure?project=PiotrFerenc_mash2&metric=duplicated_lines_density)](https://sonarcloud.io/summary/new_code?id=PiotrFerenc_mash2) [![Lines of Code](https://sonarcloud.io/api/project_badges/measure?project=PiotrFerenc_mash2&metric=ncloc)](https://sonarcloud.io/summary/new_code?id=PiotrFerenc_mash2) [![Reliability Rating](https://sonarcloud.io/api/project_badges/measure?project=PiotrFerenc_mash2&metric=reliability_rating)](https://sonarcloud.io/summary/new_code?id=PiotrFerenc_mash2) [![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=PiotrFerenc_mash2&metric=security_rating)](https://sonarcloud.io/summary/new_code?id=PiotrFerenc_mash2)  [![Technical Debt](https://sonarcloud.io/api/project_badges/measure?project=PiotrFerenc_mash2&metric=sqale_index)](https://sonarcloud.io/summary/new_code?id=PiotrFerenc_mash2) [![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=PiotrFerenc_mash2&metric=sqale_rating)](https://sonarcloud.io/summary/new_code?id=PiotrFerenc_mash2) [![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=PiotrFerenc_mash2&metric=vulnerabilities)](https://sonarcloud.io/summary/new_code?id=PiotrFerenc_mash2)

## License

