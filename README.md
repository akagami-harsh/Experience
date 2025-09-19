# Experience

This repository contains an overview of all the pull requests I have created/merged across various open source projects, highlighting my journey through prestigious programs like Google Summer of Code and Linux Foundation Mentorship.

## Table of Content

- Open Source Contributions
  - [CNCF - JaegerTracing](https://github.com/akagami-harsh/Experience/blob/main/Jaeger/README.md)
  - Kubeflow Ecosystem
    - [kubeflow - manifests](https://github.com/akagami-harsh/Experience/blob/main/kubeflow/manifests/README.md)
    - [kubeflow - pipelines](https://github.com/akagami-harsh/Experience/blob/main/kubeflow/pipelines/README.md)

- Internship Experiences
  - [LFX Mentorship - Jaeger](#lfx-mentorship)
  - [Google Summer of Code - Kubeflow](#google-summer-of-code)

## Internship Experiences

### LFX Mentorship

The Linux Foundation Mentorship (LFX) Program is a structured open source project internship where I had the opportunity to contribute to real-world, production-grade open source projects.

#### My LFX Journey

During my LFX mentorship with [Jaeger](https://github.com/jaegertracing/jaeger), I worked on Jaeger V2, a major update that rebases all backend components on top of the OpenTelemetry Collector. My primary responsibilities included:

- Integrating storage backends (Cassandra, OpenSearch, Elasticsearch, Badger, gRPC)
- Developing comprehensive integration tests for the new backends
- Collaborating with experienced open source maintainers

You can find more details about my specific contributions and experiences in the [Jaeger section](https://github.com/akagami-harsh/Experience/blob/main/Jaeger/README.md).

#### Skills Used

- Golang
- CI/CD
- Docker
- Databases
- Microservices monitoring and tracing
- Performance testing and optimization

### Google Summer of Code

Google Summer of Code (GSoC) is Google's annual program that connects university students with open source organizations for paid software development internships. I participated in GSoC 2025 with the Kubeflow organization.

#### My GSoC Journey with Kubeflow

During my GSoC internship with [Kubeflow](https://github.com/kubeflow) (March 2025 - September 2025), I focused on enhancing the security, infrastructure, and developer experience of the Kubeflow ecosystem. My key contributions included:

**🔐 Security Enhancements:**
- Designed and implemented production-grade deployment of SeaweedFS with multi-tenancy (namespace isolation) as a secure replacement for Minio
- Enhanced security by enforcing Pod Security Standards (PSS) baseline/restricted policies across all Kubeflow components including Notebooks, Katib, KServe, and Istio ingress gateway
- Implemented comprehensive security testing and validation frameworks

**🚀 Infrastructure & CI/CD:**
- Migrated Istio images from DockerHub to Google Container Registry (GCR) for improved reliability and security
- Enhanced CI/CD workflows in the kubeflow/manifests repository with automated code quality checks
- Implemented pre-commit hooks for automated linting, formatting, and validation
- Refactored shared testing components to improve maintainability and reduce code duplication

**🔧 Developer Experience:**
- Extended Kubeflow Pipelines (KFP) environment variables for better configuration management
- Improved documentation and developer onboarding processes
- Contributed to migration efforts for the kubeflow/dashboard repository
- Fixed critical issues in cluster deployment configurations

You can find detailed information about my contributions in the [Kubeflow Manifests](https://github.com/akagami-harsh/Experience/blob/main/kubeflow/manifests/README.md) and [Kubeflow Pipelines](https://github.com/akagami-harsh/Experience/blob/main/kubeflow/pipelines/README.md) sections.

#### Skills & Technologies Used

- **Cloud Native**: Kubernetes, Kustomize, Helm
- **Security**: Pod Security Standards, RBAC, Multi-tenancy
- **Storage**: SeaweedFS, Persistent Volumes, Storage Classes
- **Service Mesh**: Istio, Ingress Controllers
- **CI/CD**: GitHub Actions, Pre-commit hooks, Automated testing
- **Languages**: Go, Python, YAML
- **Tools**: Docker, kubectl, kustomize

## About This Repository

This repository serves as a portfolio of my open source contributions. Each project has its own directory with a dedicated README.md containing:

1. Project overview
2. Technologies used
3. My contributions with links to relevant PRs
4. Experiences and learnings

Feel free to explore the different projects I've contributed to!
