job "dp-search-builder" {
  datacenters = ["eu-west-1"]
  region      = "eu"
  type        = "service"

  // Make sure that this API is only run on the publishing nodes
  constraint {
    attribute = "${node.class}"
    value     = "publishing"
  }

  update {
    stagger          = "60s"
    min_healthy_time = "30s"
    healthy_deadline = "5m"
    max_parallel     = 1
    auto_revert      = true
  }

  group "publishing" {
    count = "{{PUBLISHING_TASK_COUNT}}"

    restart {
      attempts = 3
      delay    = "15s"
      interval = "1m"
      mode     = "delay"
    }

    task "dp-search-builder" {
      driver = "docker"

      artifact {
        source = "s3::https://s3-eu-west-1.amazonaws.com/{{DEPLOYMENT_BUCKET}}/dp-search-builder/{{REVISION}}.tar.gz"
      }

      config {
        command = "${NOMAD_TASK_DIR}/start-task"

        args = ["./dp-search-builder"]

        image = "{{ECR_URL}}:concourse-{{REVISION}}"

        port_map {
          http = "${NOMAD_PORT_http}"
        }
      }

      service {
        name = "dp-search-builder"
        tags = ["publishing"]
        check {
          type     = "http"
          path     = "/health"
          interval = "10s"
          timeout  = "2s"
        }
      }

      resources {
        cpu    = "{{PUBLISHING_RESOURCE_CPU}}"
        memory = "{{PUBLISHING_RESOURCE_MEM}}"

        network {
          port "http" {}
        }
      }

      template {
        source      = "${NOMAD_TASK_DIR}/vars-template"
        destination = "${NOMAD_TASK_DIR}/vars"
      }

      vault {
        policies = ["dp-search-builder"]
      }
    }
  }
}
