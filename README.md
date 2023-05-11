# container-tools

### Project structure
```javascript
.
├── README.md
├── devspace
│   ├── devspace.yaml
│   ├── devspace_start.sh
│   └── services
│       ├── publisher-backend
│       │   ├── Dockerfile
│       │   ├── devspace.yaml
│       │   ├── devspace_start.sh
│       │   ├── go.mod
│       │   ├── helm
│       │   │   ├── Chart.yaml
│       │   │   ├── charts
│       │   │   ├── templates
│       │   │   │   ├── deployment.yaml
│       │   │   │   └── service.yaml
│       │   │   └── values.yaml
│       │   └── main.go
│       └── scorer-backend
│           ├── Dockerfile
│           ├── devspace.yaml
│           ├── devspace_start.sh
│           ├── helm
│           │   ├── Chart.yaml
│           │   ├── charts
│           │   ├── templates
│           │   │   ├── deployment.yaml
│           │   │   └── service.yaml
│           │   └── values.yaml
│           ├── main.py
│           └── requirements.txt
└── skaffold
    ├── services
    │   ├── publisher-backend
    │   │   ├── Dockerfile
    │   │   ├── go.mod
    │   │   ├── helm
    │   │   │   ├── Chart.yaml
    │   │   │   ├── charts
    │   │   │   ├── templates
    │   │   │   │   ├── deployment.yaml
    │   │   │   │   └── service.yaml
    │   │   │   └── values.yaml
    │   │   ├── main.go
    │   │   └── skaffold.yaml
    │   └── scorer-backend
    │       ├── Dockerfile
    │       ├── helm
    │       │   ├── Chart.yaml
    │       │   ├── charts
    │       │   ├── templates
    │       │   │   ├── deployment.yaml
    │       │   │   └── service.yaml
    │       │   └── values.yaml
    │       ├── main.py
    │       ├── requirements.txt
    │       └── skaffold.yaml
    └── skaffold.yaml
```